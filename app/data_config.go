// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"
)

// EXISTING_CODE

var configLock atomic.Uint32

func (a *App) loadConfig(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !configLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer configLock.CompareAndSwap(1, 0)

	if !a.config.NeedsUpdate(a.forceConfig()) {
		return nil
	}

	// EXISTING_CODE
	_ = errorChan
	// EXISTING_CODE
	// EXISTING_CODE
	// do not remove
	// EXISTING_CODE

	return nil
}

func (a *App) forceConfig() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
