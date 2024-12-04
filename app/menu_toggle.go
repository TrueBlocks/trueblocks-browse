package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) ToggleHeader(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	which := "header"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleMenu(cb *menu.CallbackData) {
	which := "menu"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleHelp(cb *menu.CallbackData) {
	which := "help"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleFooter(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	which := "footer"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleAccordion(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}

	// TODO: The concepts of View and route are intermingled. In some cases `route` has
	// TODO: a leading `/` and in others it doesn't. This is confusing. Can we
	// TODO: collapse these two ideas into one? It's also intermingled with SubRoute
	// TODO: (i.e., address for history view for example) and the view's active tab.
	route := a.GetLastRoute()
	route = strings.TrimPrefix(route, "/")
	parts := strings.Split(route, "/")
	route = parts[0]
	if route == "" {
		route = "project"
	}

	a.emitMsg(messages.ToggleAccordion, &messages.MessageMsg{
		String1: route,
		String2: a.GetLastTab(route),
	})
}

func (a *App) TogglePrevTab(cb *menu.CallbackData) {
	which := "prev"
	a.emitMsg(messages.SwitchTab, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleNextTab(cb *menu.CallbackData) {
	which := "next"
	a.emitMsg(messages.SwitchTab, &messages.MessageMsg{String1: which})
}
