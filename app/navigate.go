package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
)

func (a *App) Navigate(route, subRoute string) {
	if route != "/wizard" && !a.IsConfigured() {
		route, subRoute = "/wizard", ""
	} else {
		if len(subRoute) > 0 {
			subRoute = "/" + subRoute
		}
	}

	a.session.SetRoute(route, subRoute)
	a.saveSession()

	debugMsg("Message sent", route, subRoute)
	messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg(route+subRoute))
}
