package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
)

func (a *App) isConfigured() bool {
	return a.GetSessionWizard() == wizard.Okay
}

func (a *App) GetWizardState() wizard.State {
	return a.session.Wizard.State
}

func (a *App) StepWizard(step wizard.Step) wizard.State {
	a.session.Wizard.Step(step)
	a.session.Save()
	if a.isConfigured() {
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/"))
	}
	return a.GetWizardState()
}
