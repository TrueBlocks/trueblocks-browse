package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
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
		a.emitMsg(messages.Wizard, &messages.MessageMsg{State: a.session.Wizard.State})
	}()

	a.session.Wizard.Step(step)
	a.saveSession()
	return a.getWizardState()
}

func (a *App) SetWizardState(state coreTypes.WizState) {
	a.session.Wizard.State = state
}
