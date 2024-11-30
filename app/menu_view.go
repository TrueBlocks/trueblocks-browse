// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

// EXISTING_CODE

func (a *App) ProjectView(cb *menu.CallbackData) {
	a.Navigate("/", "")
}

func (a *App) HistoryView(cb *menu.CallbackData) {
	address := a.GetSelected()
	a.Navigate("/history", address.Hex())
}

func (a *App) MonitorsView(cb *menu.CallbackData) {
	a.Navigate("/monitors", "")
}

func (a *App) NamesView(cb *menu.CallbackData) {
	a.Navigate("/names", "")
}

func (a *App) AbisView(cb *menu.CallbackData) {
	a.Navigate("/abis", "")
}

func (a *App) UnchainedView(cb *menu.CallbackData) {
	a.Navigate("/unchained", "")
}

func (a *App) SettingsView(cb *menu.CallbackData) {
	a.Navigate("/settings", "")
}

func (a *App) DaemonsView(cb *menu.CallbackData) {
	a.Navigate("/daemons", "")
}

func (a *App) WizardView(cb *menu.CallbackData) {
	if a.isConfigured() {
		a.wizard.State = types.WizWelcome
		a.emitMsg(messages.Refresh, &messages.MessageMsg{
			State: string(a.wizard.State),
			Num1:  2, // 2 is the wizard step if needed
		})
	} else {
		a.StepWizard(types.WizNext)
	}
	a.Navigate("/wizard", "")
}
