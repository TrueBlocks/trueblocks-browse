package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) IsConfigured() bool {
	return a.GetWizardState() == coreTypes.Okay
}

func (a *App) GetWizardState() coreTypes.WizState {
	return a.sessions.Wizard.State
}

func (a *App) StepWizard(step coreTypes.WizStep) coreTypes.WizState {
	defer func() {
		if a.IsConfigured() {
			a.Navigate("/", "")
		}
		messages.EmitMessage(a.ctx, messages.Wizard, &messages.MessageMsg{
			State: a.sessions.Wizard.State,
		})
	}()

	a.sessions.Wizard.Step(step)
	a.saveSession()
	return a.GetWizardState()
}
