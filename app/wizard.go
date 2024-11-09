package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) GetDeferredErrors() []types.WizError {
	var wizErrs []types.WizError
	for i, err := range a.wizard.DeferredErrors {
		wizErrs = append(wizErrs, types.WizError{Count: i, Error: err.Error()})
	}
	return wizErrs
}

func (a *App) isConfigured() bool {
	return a.wizard.State == types.WizFinished
}

func (a *App) setWizardState(state types.WizState) {
	a.wizard.State = state
}

func (a *App) addDeferredError(err error) {
	a.wizard.DeferredErrors = append(a.wizard.DeferredErrors, err)
}

func (a *App) cntDeferredErrors() int {
	return len(a.wizard.DeferredErrors)
}

var stateOrder = []types.WizState{
	types.WizWelcome,
	types.WizConfig,
	types.WizRpc,
	types.WizBlooms,
	types.WizIndex,
	types.WizFinished,
}

func (a *App) StepWizard(step types.WizStep) types.WizState {
	defer func() {
		if a.isConfigured() {
			a.Navigate("/", "")
		}
		a.emitMsg(messages.Refresh, &messages.MessageMsg{
			State: string(a.wizard.State),
			Num1:  2, // 2 is the wizard step if needed
		})
	}()

	switch step {
	case types.WizFirst:
		a.wizard.State = types.WizWelcome

	case types.WizPrevious:
		for i := range stateOrder {
			if stateOrder[i] == a.wizard.State && i > 0 {
				a.wizard.State = stateOrder[i-1]
				break
			}
		}

	case types.WizNext:
		for i := range stateOrder {
			if stateOrder[i] == a.wizard.State && i < len(stateOrder)-1 {
				a.wizard.State = stateOrder[i+1]
				break
			}
		}

	case types.WizFinish:
		a.wizard.State = types.WizFinished
	}

	a.saveSession()
	logger.InfoBB("Wizard state:", a.wizard.State)

	return a.wizard.State
}
