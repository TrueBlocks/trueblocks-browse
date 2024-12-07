package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) Navigate(route, address string) {
	sep := ""
	if len(address) > 0 {
		sep = "/"
	}

	if route != "/wizard" && !a.isConfigured() {
		route, address, sep = "/wizard", "", ""
	}

	a.SetLastRoute(route, address)

	logger.Info("Message sent", route, address)
	a.emitNavigateMsg(route + sep + address)
}
