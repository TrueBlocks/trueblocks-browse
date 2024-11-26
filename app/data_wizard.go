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
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
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
	logger.InfoBY("Updating needed for wizard...")

	opts := WizardOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	opts.Chain = a.getChain()
	// EXISTING_CODE
	if wizard, meta, err := opts.WizardList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (wizard == nil) || (len(wizard) == 0) {
		// expected outcome
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.wizard = types.NewWizardContainer(opts.Chain, wizard)
		// EXISTING_CODE
		a.Navigate("/wizard", "")
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "wizard")
	}

	return nil
}

// EXISTING_CODE
type WizardOptions struct {
	Globals sdk.Globals
	Chain   string
}

func (opts *WizardOptions) WizardList() ([]types.WizError, *coreTypes.MetaData, error) {
	meta, err := sdk.GetMetaData(namesChain)
	// TODO: We've been called to check status, do wizard checks here
	return []types.WizError{}, meta, err
}

// EXISTING_CODE
