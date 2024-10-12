package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) ViewProject(cd *menu.CallbackData) {
	logger.Info("ViewProject")
	a.Navigate("/", "")
}

func (a *App) ViewHistory(cd *menu.CallbackData) {
	address := a.GetLastAddress()
	logger.Info("ViewHistory")
	a.Navigate("/history", address.Hex())
}

func (a *App) ViewMonitors(cd *menu.CallbackData) {
	logger.Info("ViewMonitors")
	a.Navigate("/monitors", "")
}

func (a *App) ViewNames(cd *menu.CallbackData) {
	logger.Info("ViewNames")
	a.Navigate("/names", "")
}

func (a *App) ViewIndexes(cd *menu.CallbackData) {
	logger.Info("ViewIndexes")
	a.Navigate("/indexes", "")
}

func (a *App) ViewManifest(cd *menu.CallbackData) {
	logger.Info("ViewManifest")
	a.Navigate("/manifest", "")
}

func (a *App) ViewAbis(cd *menu.CallbackData) {
	logger.Info("ViewAbis")
	a.Navigate("/abis", "")
}

func (a *App) ViewStatus(cd *menu.CallbackData) {
	logger.Info("ViewStatus")
	a.Navigate("/status", "")
}

func (a *App) ViewDaemons(cd *menu.CallbackData) {
	logger.Info("ViewDaemons")
	a.Navigate("/daemons", "")
}

func (a *App) ViewSettings(cd *menu.CallbackData) {
	logger.Info("ViewSettings")
	a.Navigate("/settings", "")
}

func (a *App) ViewWizard(cd *menu.CallbackData) {
	a.StepWizard(wizard.Reset)
	logger.Info("ViewWizard")
	a.Navigate("/wizard", "")
}
