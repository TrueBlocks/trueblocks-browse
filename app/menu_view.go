package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) ProjectView(cd *menu.CallbackData) {
	a.Navigate("/", "")
}

func (a *App) HistoryView(cd *menu.CallbackData) {
	address := a.GetAddress()
	a.Navigate("/history", address.Hex())
}

func (a *App) MonitorsView(cd *menu.CallbackData) {
	a.Navigate("/monitors", "")
}

func (a *App) NamesView(cd *menu.CallbackData) {
	a.Navigate("/names", "")
}

func (a *App) AbisView(cd *menu.CallbackData) {
	a.Navigate("/abis", "")
}

func (a *App) IndexesView(cd *menu.CallbackData) {
	a.Navigate("/indexes", "")
}

func (a *App) ManifestsView(cd *menu.CallbackData) {
	a.Navigate("/manifests", "")
}

func (a *App) StatusView(cd *menu.CallbackData) {
	a.Navigate("/status", "")
}

func (a *App) DaemonsView(cd *menu.CallbackData) {
	a.Navigate("/daemons", "")
}

func (a *App) SettingsView(cd *menu.CallbackData) {
	a.Navigate("/settings", "")
}

func (a *App) WizardView(cd *menu.CallbackData) {
	if a.IsConfigured() {
		a.StepWizard(wizard.Reset)
	} else {
		a.StepWizard(wizard.Next)
	}
	a.Navigate("/wizard", "")
}
