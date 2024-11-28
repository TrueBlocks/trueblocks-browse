package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

func (a *App) isConfigured() bool {
	return a.wizard.State == types.WizFinished
}

func (a *App) setWizardState(state types.WizState) {
	a.wizard.State = state
}

func (a *App) addWizErr(reason string, state types.WizState, err error) {
	wizError := types.WizError{
		Index:  a.cntWizErrs() + 1,
		State:  state,
		Reason: reason,
		Error:  err.Error(),
	}
	a.wizard.Items = append(a.wizard.Items, wizError)
}

func (a *App) cntWizErrs() int {
	return len(a.wizard.Items)
}

func (a *App) emitWizErrs() {
	for _, wizErr := range a.wizard.Items {
		a.emitErrorMsg(wizErr.ToErr(), nil)
	}
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
			State: a.wizard.State.String(),
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
	return a.wizard.State
}

const (
	WizReasonNoSession           = "could not load session file"
	WizReasonNoConfig            = "could not load config file"
	WizReasonChainNotConfigured  = "chain is not configured"
	WizReasonFailedRpcPing       = "could not connect to Rpc"
	WizReasonFailedNamesLoad     = "could not load names"
	WizReasonFailedPrepareWindow = "could not prepare window"
	WizReasonNoFreshenDaemon     = "could not start freshen daemon"
	WizReasonNoScraperDaemon     = "could not start scraper daemon"
	WizReasonNoIpfsDaemon        = "could not start Ipfs daemon"
)
