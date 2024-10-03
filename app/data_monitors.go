package app

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) MonitorPage(first, pageSize int) *types.MonitorContainer {
	first = base.Max(0, base.Min(first, len(a.monitors.Items)-1))
	last := base.Min(len(a.monitors.Items), first+pageSize)
	copy, _ := a.monitors.ShallowCopy().(*types.MonitorContainer)
	copy.Items = a.monitors.Items[first:last]
	return copy
}

var monitorLock atomic.Uint32

func (a *App) loadMonitors(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !monitorLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer monitorLock.CompareAndSwap(1, 0)

	if !a.monitors.NeedsUpdate() {
		return nil
	}

	chain := a.globals.Chain
	opts := sdk.MonitorsOptions{
		Globals: sdk.Globals{
			Verbose: true,
			Chain:   chain,
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
		a.monitors = types.NewMonitorContainer(chain)
		for _, mon := range monitors {
			mon.Name = a.names.NamesMap[mon.Address].Name
			a.monitors.Items = append(a.monitors.Items, mon)
			a.monitors.MonitorMap[mon.Address] = mon
		}
		// TODO: Use core's sorting mechanism (see SortChunkStats for example)
		sort.Slice(a.monitors.Items, func(i, j int) bool {
			return a.monitors.Items[i].NRecords < a.monitors.Items[j].NRecords
		})
		a.monitors.Summarize()
		messages.SendInfo(a.ctx, "Loaded monitors")
	}
	return nil
}
