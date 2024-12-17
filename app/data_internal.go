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
)

// EXISTING_CODE

var internalsLock atomic.Uint32

func (a *App) loadInternals(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadInternals", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !internalsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer internalsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.internals.NeedsUpdate() {
		return nil
	}
	updater := a.internals.Updater
	defer func() {
		a.internals.Updater = updater
	}()
	logger.InfoBY("Updating internals...")

	if items, meta, err := a.pullInternals(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no internals found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.internals = types.NewInternalContainer(a.getChain(), items, a.getLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.internals.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "internals")
	}

	return nil
}

func (a *App) pullInternals() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
