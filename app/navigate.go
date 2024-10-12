package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) Navigate(route, subRoute string) {
	if route != "/wizard" && !a.isConfigured() {
		return
	}

	logger.Info("Message sent", route, subRoute)
	if len(subRoute) > 0 {
		subRoute = "/" + subRoute
	}
	a.session.SetRoute(route, subRoute)
	a.saveSession()
	messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg(route+subRoute))
}
