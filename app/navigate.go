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
		route, subRoute, sep = "/wizard", "", ""
	}

	a.SetRoute(route, subRoute)

	debugMsg("Message sent", route, subRoute)
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: route + sep + subRoute,
	})
}
