// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var daemonsLock atomic.Uint32

func (a *App) loadDaemons(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadDaemons", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !daemonsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer daemonsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.daemons.NeedsUpdate() {
		return nil
	}
	updater := a.daemons.Updater
	defer func() {
		a.daemons.Updater = updater
	}()
	logger.InfoBY("Updating daemons...")

	if items, meta, err := a.pullDaemons(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		// this outcome is okay
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.daemons = types.NewDaemonContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.daemons.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "daemons")
	}

	return nil
}

func (a *App) pullDaemons() (items []types.Daemon, meta *types.Meta, err error) {
	// EXISTING_CODE
	meta, err = sdk.GetMetaData(namesChain)
	// TODO: We've been called to update the status of the daemons. Do so here.
	return []types.Daemon{}, meta, err
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
