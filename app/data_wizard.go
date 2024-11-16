// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"
)

// EXISTING_CODE

var wizardLock atomic.Uint32

func (a *App) loadWizard(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !wizardLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer wizardLock.CompareAndSwap(1, 0)

	if !a.wizard.NeedsUpdate(a.forceWizard()) {
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
	a.emitInfoMsg("Loaded wizard", "")

	return nil
}

func (a *App) forceWizard() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
