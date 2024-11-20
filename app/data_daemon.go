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

// -------------------------------------------------------------------
func (a *App) loadDaemons(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadDaemons", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !daemonLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer daemonLock.CompareAndSwap(1, 0)

	if !a.daemons.NeedsUpdate(a.forceDaemon()) {
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
	a.emitInfoMsg("Loaded daemons", "")

	return nil
}

// -------------------------------------------------------------------
func (a *App) forceDaemon() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
