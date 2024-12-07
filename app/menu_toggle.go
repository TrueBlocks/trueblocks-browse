package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) ToggleAppHeader(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	which := "header"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleAppMenu(cb *menu.CallbackData) {
	which := "menu"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleAppHelp(cb *menu.CallbackData) {
	which := "help"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleAppFooter(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	which := "footer"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) ToggleViewHeader(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}

	route, _ := a.session.GetRouteAndSub()
	route = strings.Trim(route, "/")
	tab := a.GetLastTab(route)
	newState := a.ToggleHeader(route, tab)

	a.emitMsg(messages.ToggleHeader, &messages.MessageMsg{
		String1: route,
		String2: tab,
		Bool:    newState,
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
