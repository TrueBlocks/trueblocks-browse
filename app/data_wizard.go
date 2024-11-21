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

var wizardLock atomic.Uint32

func (a *App) loadWizard(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadWizard", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !wizardLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer wizardLock.CompareAndSwap(1, 0)

	if !a.wizard.NeedsUpdate() {
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
	a.emitLoadingMsg(messages.Loaded, "wizard")

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
