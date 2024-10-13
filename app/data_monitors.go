package app

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var monitorMutex sync.Mutex

// Find: NewViews
func (a *App) MonitorPage(first, pageSize int) *types.MonitorContainer {
	first = base.Max(0, base.Min(first, len(a.monitors.Monitors)-1))
	last := base.Min(len(a.monitors.Monitors), first+pageSize)
	copy, _ := a.monitors.ShallowCopy().(*types.MonitorContainer)
	copy.Monitors = a.monitors.Monitors[first:last]
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

	if !a.monitors.NeedsUpdate(a.nameChange()) {
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
			a.monitors.Monitors = append(a.monitors.Monitors, mon)
		}
		// TODO: Use core's sorting mechanism (see SortChunkStats for example)
		sort.Slice(a.monitors.Monitors, func(i, j int) bool {
			return a.monitors.Monitors[i].NRecords < a.monitors.Monitors[j].NRecords
		})
		a.monitors.Summarize()
		messages.EmitInfo(a.ctx, "Loaded monitors")
	}
	return nil
}

func (a *App) ModifyMonitors(modData *ModifyData) error {
	if !monitorLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer monitorLock.CompareAndSwap(1, 0)

	opFromString := func(op string) crud.Operation {
		m := map[string]crud.Operation{
			"delete":   crud.Delete,
			"undelete": crud.Undelete,
			"remove":   crud.Remove,
		}
		return m[op]
	}

	op := opFromString(modData.Operation)
	opts := sdk.MonitorsOptions{
		Addrs:    []string{modData.Address.Hex()},
		Delete:   op == crud.Delete,
		Undelete: op == crud.Undelete,
		Remove:   op == crud.Remove,
		Globals:  a.globals,
	}

	if _, _, err := opts.Monitors(); err != nil {
		messages.EmitError(a.ctx, err)
		return err
	}

	newArray := []coreTypes.Monitor{}
	for _, mon := range a.monitors.Monitors {
		if mon.Address == modData.Address {
			switch op {
			case crud.Delete:
				mon.Deleted = true
			case crud.Undelete:
				mon.Deleted = false
			case crud.Remove:
				continue
			}
		}
		newArray = append(newArray, mon)
	}
	monitorMutex.Lock()
	defer monitorMutex.Unlock()

	a.monitors.Monitors = newArray
	a.monitors.Summarize()
	return nil
}
