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

func (a *App) loadConfig(wg *sync.WaitGroup, errorChan chan error) {
	_ = errorChan
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !configLock.CompareAndSwap(0, 1) {
		return
	}
	defer configLock.CompareAndSwap(1, 0)

	if !a.config.NeedsUpdate(a.forceAbi()) {
		return
	}
}
