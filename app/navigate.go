package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) Navigate(route, subRoute string) {
	sep := ""
	if len(subRoute) > 0 {
		sep = "/"
	}

	if route != "/wizard" && !a.isConfigured() {
		route, subRoute, sep = "/wizard", "", ""
	}

	a.SetRoute(route, subRoute)

	logger.Info("Message sent", route, subRoute)
	a.emitNavigateMsg(route + sep + subRoute)
}
