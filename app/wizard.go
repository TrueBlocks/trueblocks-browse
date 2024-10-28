package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

func (a *App) IsConfigured() bool {
	return a.GetWizardState() == types.Okay
}

func (a *App) GetWizardState() types.State {
	return a.sessions.Wizard.State
}

func (a *App) StepWizard(step types.Step) types.State {
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
