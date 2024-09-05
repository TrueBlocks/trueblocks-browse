package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

// Find: NewViews
func (a *App) ViewPortfolio(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewPortfolio")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/"))
		a.SetLast("route", "/")
	}
}

func (a *App) ViewHistory(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewHistory")
		subRoute := a.GetLastSub("/history")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/history"+subRoute))
		a.SetLast("route", "/history"+subRoute)
	}
}

func (a *App) ViewMonitors(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewMonitors")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/monitors"))
		a.SetLast("route", "/monitors")
	}
}

func (a *App) ViewNames(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewNames")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/names"))
		a.SetLast("route", "/names")
	}
}

func (a *App) ViewIndexes(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewIndexes")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/indexes"))
		a.SetLast("route", "/indexes")
	}
}

func (a *App) ViewManifest(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewManifest")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/manifest"))
		a.SetLast("route", "/manifest")
	}
}

func (a *App) ViewAbis(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewAbis")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/abis"))
		a.SetLast("route", "/abis")
	}
}

func (a *App) ViewStatus(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewStatus")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/status"))
		a.SetLast("route", "/status")
	}
}

func (a *App) ViewDaemons(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewDaemons")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/daemons"))
		a.SetLast("route", "/daemons")
	}
}

func (a *App) ViewSettings(cd *menu.CallbackData) {
	if a.isConfigured() {
		logger.Info("ViewSettings")
		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/settings"))
		a.SetLast("route", "/settings")
	}
}

func (a *App) ViewWizard(cd *menu.CallbackData) {
	a.StepWizard(wizard.Reset)
	logger.Info("ViewWizard")
	messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/wizard"))
	a.SetLast("route", "/wizard")
}
