package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
)

func (a *App) Navigate(route, subRoute string) {
	sep := ""
	if len(subRoute) > 0 {
		sep = "/"
	}

	if route != "/wizard" && !a.IsConfigured() {
		route, subRoute = "/wizard", ""
	}

	a.session.SetRoute(route, subRoute)
	a.saveSession()

	debugMsg("Message sent", route, subRoute)
	messages.EmitNavigate(a.ctx, route+sep+subRoute)
}
