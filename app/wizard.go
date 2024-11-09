package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) getWizardState() coreTypes.WizState {
	return a.session.Wizard
}

func (a *App) isConfigured() bool {
	return a.getWizardState() == coreTypes.WizFinished
}

func (a *App) SetWizardState(state coreTypes.WizState) {
	a.session.Wizard = state
}

func (a *App) GetDeferredErrors() []types.WizardError {
	var wizErrs []types.WizardError
	for i, err := range a.wizard.DeferredErrors {
		wizErrs = append(wizErrs, types.WizardError{Count: i, Error: err.Error()})
	}
	return wizErrs
}

func (a *App) addDeferredError(err error) {
	a.wizard.DeferredErrors = append(a.wizard.DeferredErrors, err)
}

func (a *App) cntDeferredErrors() int {
	return len(a.wizard.DeferredErrors)
}

var stateOrder = []coreTypes.WizState{
	coreTypes.WizWelcome,
	coreTypes.WizError,
	coreTypes.WizConfig,
	coreTypes.WizRpc,
	coreTypes.WizBlooms,
	coreTypes.WizIndex,
	coreTypes.WizFinished,
}

func (a *App) StepWizard(step coreTypes.WizStep) coreTypes.WizState {
	defer func() {
		// TODO: Remove this
		if step != coreTypes.WizFirst && len(a.wizard.DeferredErrors) > 0 {
			a.wizard.DeferredErrors = a.wizard.DeferredErrors[:len(a.wizard.DeferredErrors)-1]
		}

		if a.isConfigured() {
			a.Navigate("/", "")
		}
		a.emitMsg(messages.Refresh, &messages.MessageMsg{
			State: a.session.Wizard,
			Num1:  2, // 2 is the wizard step if needed
		})
	}()

	switch step {
	case coreTypes.WizReset:
		a.session.Wizard = coreTypes.WizError
	case coreTypes.WizFirst:
		a.session.Wizard = coreTypes.WizWelcome
	case coreTypes.WizPrevious:
		if a.session.Wizard == coreTypes.WizConfig {
			a.session.Wizard = coreTypes.WizWelcome
		} else {
			for i := range stateOrder {
				if stateOrder[i] == a.session.Wizard && i > 0 {
					a.session.Wizard = stateOrder[i-1]
					break
				}
			}
		}
	case coreTypes.WizNext:
		if a.session.Wizard == coreTypes.WizWelcome {
			a.session.Wizard = coreTypes.WizConfig
		} else {
			for i := range stateOrder {
				if stateOrder[i] == a.session.Wizard && i < len(stateOrder)-1 {
					a.session.Wizard = stateOrder[i+1]
					break
				}
			}
		}
	case coreTypes.WizFinish:
		a.session.Wizard = coreTypes.WizFinished
	}

	a.saveSession()
	return a.getWizardState()
}
