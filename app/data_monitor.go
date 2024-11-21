// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var monitorMutex sync.Mutex

// EXISTING_CODE

var monitorLock atomic.Uint32

func (a *App) loadMonitors(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadMonitors", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !monitorLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer monitorLock.CompareAndSwap(1, 0)

	if !a.monitors.NeedsUpdate(a.forceMonitor()) {
		return nil
	}

	opts := sdk.MonitorsOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	// EXISTING_CODE
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
		// EXISTING_CODE
		for i := 0; i < len(monitors); i++ {
			monitors[i].Name = a.namesMap[monitors[i].Address].Name
		}
		// EXISTING_CODE
		a.meta = *meta
		a.monitors = types.NewMonitorContainer(opts.Chain, monitors)
		// EXISTING_CODE
		// TODO: Use core's sorting mechanism (see SortChunkStats for example)
		sort.Slice(a.monitors.Items, func(i, j int) bool {
			if a.monitors.Items[i].NRecords == a.monitors.Items[j].NRecords {
				return a.monitors.Items[i].Address.Hex() < a.monitors.Items[j].Address.Hex()
			}
			return a.monitors.Items[i].NRecords < a.monitors.Items[j].NRecords
		})
		// EXISTING_CODE
		a.emitInfoMsg("Loaded monitors", "")
	}

	return nil
}

func (a *App) forceMonitor() (force bool) {
	// EXISTING_CODE
	force = a.forceName()
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (a *App) ModifyMonitor(modData *ModifyData) error {
	if !monitorLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer monitorLock.CompareAndSwap(1, 0)

	op := crud.OpFromString(modData.Operation)
	opts := sdk.MonitorsOptions{
		Addrs:    []string{modData.Address.Hex()},
		Delete:   op == crud.Delete,
		Undelete: op == crud.Undelete,
		Remove:   op == crud.Remove,
		Globals:  a.getGlobals(false /* verbose */),
	}

	if _, _, err := opts.Monitors(); err != nil {
		a.emitErrorMsg(err, nil)
		return err
	}

	newArray := []coreTypes.Monitor{}
	for _, mon := range a.monitors.Items {
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

	a.monitors.Items = newArray
	return nil
}

// EXISTING_CODE
