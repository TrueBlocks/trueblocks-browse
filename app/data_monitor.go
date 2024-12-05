// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var monitorMutex sync.Mutex

// EXISTING_CODE

var monitorsLock atomic.Uint32

func (a *App) loadMonitors(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadMonitors", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !monitorsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer monitorsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.monitors.NeedsUpdate() {
		return nil
	}
	updater := a.monitors.Updater
	defer func() {
		a.monitors.Updater = updater
	}()
	logger.InfoBY("Updating monitors...")

	if items, meta, err := a.pullMonitors(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no monitors found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		for i := 0; i < len(items); i++ {
			items[i].Name = a.namesMap[items[i].Address].Name
		}
		// EXISTING_CODE
		a.meta = *meta
		a.monitors = types.NewMonitorContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.monitors.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "monitors")
	}

	return nil
}

func (a *App) pullMonitors() (items []types.Monitor, meta *types.Meta, err error) {
	// EXISTING_CODE
	opts := sdk.MonitorsOptions{
		Globals: sdk.Globals{
			Chain:   a.getChain(),
			Verbose: true,
		},
	}
	return opts.MonitorsList()
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
