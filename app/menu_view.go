// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

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

func (a *App) IndexesView(cb *menu.CallbackData) {
	a.Navigate("/indexes", "")
}

func (a *App) ManifestsView(cb *menu.CallbackData) {
	a.Navigate("/manifests", "")
}

func (a *App) StatusView(cb *menu.CallbackData) {
	a.Navigate("/status", "")
}

func (a *App) SettingsView(cb *menu.CallbackData) {
	a.Navigate("/settings", "")
}

func (a *App) DaemonsView(cb *menu.CallbackData) {
	a.Navigate("/daemons", "")
}

func (a *App) SessionView(cb *menu.CallbackData) {
	a.Navigate("/session", "")
}

func (a *App) ConfigView(cb *menu.CallbackData) {
	a.Navigate("/config", "")
}

func (a *App) WizardView(cb *menu.CallbackData) {
	if a.isConfigured() {
		a.StepWizard(coreTypes.WizReset)
	} else {
		a.StepWizard(coreTypes.WizNext)
	}
	a.Navigate("/wizard", "")
}
