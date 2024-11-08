package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) isConfigured() bool {
	return a.getWizardState() == coreTypes.Okay
}

func (a *App) getWizardState() coreTypes.WizState {
	return a.session.Wizard.State
}

func (a *App) StepWizard(step coreTypes.WizStep) coreTypes.WizState {
	defer func() {
		if a.isConfigured() {
			a.Navigate("/", "")
		}
		a.emitMsg(messages.Refresh, &messages.MessageMsg{State: a.session.Wizard.State})
	}()

	a.session.Wizard.Step(step)
	a.saveSession()
	return a.getWizardState()
}

func (a *App) SetWizardState(state coreTypes.WizState) {
	a.session.Wizard.State = state
}

func (a *App) cntDeferredErrors() int {
	return len(a.wizard.DeferredErrors)
}

func (a *App) addDeferredError(err error) {
	a.wizard.DeferredErrors = append(a.wizard.DeferredErrors, err)
}

func (a *App) GetDeferredErrors() []types.WizardError {
	var wizErrs []types.WizardError
	for i, err := range a.wizard.DeferredErrors {
		wizErrs = append(wizErrs, types.WizardError{Count: i, Error: err.Error()})
	}
	return wizErrs
}
