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

var configLock atomic.Uint32

func (a *App) loadConfig(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadConfig", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !configLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer configLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.config.NeedsUpdate() {
		return nil
	}
	updater := a.config.Updater
	defer func() {
		a.config.Updater = updater
	}()
	logger.InfoBY("Updating needed for config...")

	// EXISTING_CODE
	_ = errorChan
	// EXISTING_CODE
	// EXISTING_CODE
	// do not remove
	// EXISTING_CODE
	// EXISTING_CODE
	// do not remove
	// EXISTING_CODE
	a.emitLoadingMsg(messages.Loaded, "config")

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
