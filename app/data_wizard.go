// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	// TOOD: BOGUS "github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

var wizardLock atomic.Uint32

func (a *App) loadWizard(wg *sync.WaitGroup, errorChan chan error) {
	_ = errorChan
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !wizardLock.CompareAndSwap(0, 1) {
		return
	}
	defer wizardLock.CompareAndSwap(1, 0)

	if !a.wizard.NeedsUpdate(false) {
		return
	}

	// TOOD: BOGUS a.wizard = types.NewWizardContainer(a.session.LastChain, a.wizard.Items)
	// TOOD: BOGUS a.wizard.Summarize()
}
