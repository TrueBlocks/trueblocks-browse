// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"
)

// EXISTING_CODE

var daemonLock atomic.Uint32

func (a *App) loadDaemon(wg *sync.WaitGroup, errorChan chan error) {
	_ = errorChan
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !daemonLock.CompareAndSwap(0, 1) {
		return
	}
	defer daemonLock.CompareAndSwap(1, 0)

	if !a.daemons.NeedsUpdate(false) {
		return
	}
}
