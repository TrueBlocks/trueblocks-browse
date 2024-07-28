package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Find: NewViews
func (a *App) ViewHome(cd *menu.CallbackData) {
	logger.Info("View Home")
	runtime.EventsEmit(a.ctx, "navigate", "/")
}

func (a *App) ViewHistory(cd *menu.CallbackData) {
	logger.Info("View History")
	last := a.GetLast("address")
	if len(last) == 0 {
		last = "trueblocks.eth"
	}
	runtime.EventsEmit(a.ctx, "navigate", "/history/"+last)
}

func (a *App) ViewMonitors(cd *menu.CallbackData) {
	logger.Info("View Monitors")
	runtime.EventsEmit(a.ctx, "navigate", "/monitors")
}

func (a *App) ViewNames(cd *menu.CallbackData) {
	logger.Info("View Names")
	runtime.EventsEmit(a.ctx, "navigate", "/names")
}

func (a *App) ViewServers(cd *menu.CallbackData) {
	logger.Info("View Servers")
	runtime.EventsEmit(a.ctx, "navigate", "/servers")
}

func (a *App) ViewSettings(cd *menu.CallbackData) {
	logger.Info("View Settings")
	runtime.EventsEmit(a.ctx, "navigate", "/settings")
}
