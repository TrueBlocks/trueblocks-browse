// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

// EXISTING_CODE

func (a *App) ProjectView(cb *menu.CallbackData) {
	a.Navigate("project")
}

func (a *App) HistoryView(cb *menu.CallbackData) {
	if strings.Contains(a.GetLastRoute(), "history") {
		a.ToggleNextTab(cb)
	} else {
		a.Navigate("history")
	}
}

func (a *App) MonitorsView(cb *menu.CallbackData) {
	a.Navigate("monitors")
}

func (a *App) SharingView(cb *menu.CallbackData) {
	if strings.Contains(a.GetLastRoute(), "sharing") {
		a.ToggleNextTab(cb)
	} else {
		a.Navigate("sharing")
	}
}

func (a *App) UnchainedView(cb *menu.CallbackData) {
	if strings.Contains(a.GetLastRoute(), "unchained") {
		a.ToggleNextTab(cb)
	} else {
		a.Navigate("unchained")
	}
}

func (a *App) SettingsView(cb *menu.CallbackData) {
	if strings.Contains(a.GetLastRoute(), "settings") {
		a.ToggleNextTab(cb)
	} else {
		a.Navigate("settings")
	}
}

func (a *App) DaemonsView(cb *menu.CallbackData) {
	a.Navigate("daemons")
}

func (a *App) WizardView(cb *menu.CallbackData) {
	if a.isConfigured() {
		a.setWizState(types.WizWelcome)
		a.emitMsg(messages.Refresh, &messages.MessageMsg{
			State: string(a.getWizState()),
			Num1:  2, // 2 is the wizard step if needed
		})
	} else {
		a.StepWizard(types.WizNext)
	}
	a.Navigate("wizard")
}
