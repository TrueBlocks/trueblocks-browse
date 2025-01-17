// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
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

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.wizard.NeedsUpdate() {
		return nil
	}
	updater := a.wizard.Updater
	defer func() {
		a.wizard.Updater = updater
	}()
	logger.InfoBY("Updating wizard...")

	if items, meta, err := a.pullWizards(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		// this outcome is okay
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.wizard = types.NewWizardContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.wizard.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "wizard")
	}

	return nil
}

func (a *App) pullWizards() (items []types.WizError, meta *types.Meta, err error) {
	// EXISTING_CODE
	meta, err = sdk.GetMetaData(namesChain)
	// TODO: We've been called to check status, do wizard checks here
	return []types.WizError{}, meta, err
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
