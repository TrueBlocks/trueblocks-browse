// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
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

	if !a.config.NeedsUpdate() {
		return nil
	}

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
