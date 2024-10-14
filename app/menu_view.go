package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) ViewProject(cd *menu.CallbackData) {
	debugMsg("ViewProject")
	a.Navigate("/", "")
}

func (a *App) ViewHistory(cd *menu.CallbackData) {
	debugMsg("ViewHistory")

	address := a.GetAddress()
	a.Navigate("/history", address.Hex())
}

func (a *App) ViewMonitors(cd *menu.CallbackData) {
	debugMsg("ViewMonitors")
	a.Navigate("/monitors", "")
}

func (a *App) ViewNames(cd *menu.CallbackData) {
	debugMsg("ViewNames")
	a.Navigate("/names", "")
}

func (a *App) ViewIndexes(cd *menu.CallbackData) {
	debugMsg("ViewIndexes")
	a.Navigate("/indexes", "")
}

func (a *App) ViewManifest(cd *menu.CallbackData) {
	debugMsg("ViewManifest")
	a.Navigate("/manifest", "")
}

func (a *App) ViewAbis(cd *menu.CallbackData) {
	debugMsg("ViewAbis")
	a.Navigate("/abis", "")
}

func (a *App) ViewStatus(cd *menu.CallbackData) {
	debugMsg("ViewStatus")
	a.Navigate("/status", "")
}

func (a *App) ViewDaemons(cd *menu.CallbackData) {
	debugMsg("ViewDaemons")
	a.Navigate("/daemons", "")
}

func (a *App) ViewSettings(cd *menu.CallbackData) {
	debugMsg("ViewSettings")
	a.Navigate("/settings", "")
}

func (a *App) ViewWizard(cd *menu.CallbackData) {
	debugMsg("ViewWizard")

	if a.IsConfigured() {
		a.StepWizard(wizard.Reset)
	} else {
		a.StepWizard(wizard.Next)
	}
	a.Navigate("/wizard", "")
}
