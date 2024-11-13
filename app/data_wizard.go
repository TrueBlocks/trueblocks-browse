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
	_ = errorChan
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

	// opts := sdk.WizardOptions{
	// 	Globals: a.getGlobals(),
	// }
	// // EXISTING_CODE
	// // EXISTING_CODE
	// opts.Verbose = true

	// if wizard, meta, err := opts.WizardList(); err != nil {
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else if (wizard == nil) || (len(wizard) == 0) {
	// 	err = fmt.Errorf("no wizard found")
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else {
	// 	// EXISTING_CODE
	// 	// EXISTING_CODE
	// 	a.meta = *meta
	// 	a.wizard = types.NewWizardContainer(opts.Chain, wizard)
	// 	// EXISTING_CODE
	// 	// EXISTING_CODE
	// 	a.emitInfoMsg("Loaded wizard", "")
	// }

	return nil
}

func (a *App) forceWizard() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
