// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
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

	if !a.daemons.NeedsUpdate() {
		return nil
	}
	updater := a.daemons.Updater
	defer func() {
		a.daemons.Updater = updater
	}()
	logger.InfoBY("Updating needed for daemons...")

	// EXISTING_CODE
	_ = errorChan
	// EXISTING_CODE
	// EXISTING_CODE
	// EXISTING_CODE
	// EXISTING_CODE
	// EXISTING_CODE
	a.emitLoadingMsg(messages.Loaded, "daemons")

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
