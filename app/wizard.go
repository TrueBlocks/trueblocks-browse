package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
)

func (a *App) IsConfigured() bool {
	return a.GetWizardState() == wizard.Okay
}

func (a *App) GetWizardState() wizard.State {
	return a.session.Wizard.State
}

func (a *App) StepWizard(step wizard.Step) wizard.State {
	defer func() {
		if a.IsConfigured() {
			a.Navigate("/", "")
		}
		messages.EmitMessage(a.ctx, messages.Wizard, &messages.MessageMsg{
			State: a.session.Wizard.State,
		})
	}()

	a.session.Wizard.Step(step)
	a.saveSession()
	return a.GetWizardState()
}
