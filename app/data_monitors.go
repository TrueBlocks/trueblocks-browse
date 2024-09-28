package app

import (
	"fmt"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) MonitorPage(first, pageSize int) *types.MonitorContainer {
	first = base.Max(0, base.Min(first, len(a.monitors.Items)-1))
	last := base.Min(len(a.monitors.Items), first+pageSize)
	copy := a.monitors.ShallowCopy().(*types.MonitorContainer)
	copy.Items = a.monitors.Items[first:last]
	return copy
}

func (a *App) loadMonitors(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	opts := sdk.MonitorsOptions{
		Staged: true,
		Globals: sdk.Globals{
			Verbose: true,
			Chain:   a.globals.Chain,
		},
	}

	if monitors, meta, err := opts.MonitorsList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (monitors == nil) || (len(monitors) == 0) {
		err = fmt.Errorf("no monitors found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		if len(a.monitors.Items) == len(monitors) {
			return nil
		}
		a.monitors = types.MonitorContainer{}
		a.monitors.MonitorMap = make(map[base.Address]coreTypes.Monitor)
		for _, mon := range monitors {
			mon.Name = a.names.NamesMap[mon.Address].Name
			a.monitors.Items = append(a.monitors.Items, mon)
			a.monitors.MonitorMap[mon.Address] = mon
		}
		sort.Slice(a.monitors.Items, func(i, j int) bool {
			return a.monitors.Items[i].NRecords < a.monitors.Items[j].NRecords
		})
		a.monitors.Summarize()
		messages.SendInfo(a.ctx, "Loaded monitors")
	}
	return nil
}
