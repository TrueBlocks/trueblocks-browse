package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
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

	tabs := a.getTabs()
	route := a.getLastRoute()
	tab := a.getLastTab(route)
	for _, t := range tabs {
		if t == tab {
			isOn := a.GetHeaderOn(route, tab)
			a.SetHeaderOn(route, tab, !isOn)
			a.emitMsg(messages.ToggleHeader, &messages.MessageMsg{
				String1: route,
				String2: tab,
				Bool:    !isOn,
			})
		}
	}
}

func (a *App) TogglePrevTab(cb *menu.CallbackData) {
	tabs := a.getTabs()
	route := a.getLastRoute()
	tab := a.getLastTab(route)
	for i, t := range tabs {
		if t == tab {
			newTab := tabs[(i-1+len(tabs))%len(tabs)]
			a.setLastTab(route, newTab)
			a.emitNavigateMsg(route, newTab)
			return
		}
	}
	logger.Error("TogglePrevTab: should never happen")
	// should never happen
}

func (a *App) ToggleNextTab(cb *menu.CallbackData) {
	tabs := a.getTabs()
	route := a.getLastRoute()
	tab := a.getLastTab(route)
	for i, t := range tabs {
		if t == tab {
			newTab := tabs[(i+1)%len(tabs)]
			a.setLastTab(route, newTab)
			a.emitNavigateMsg(route, newTab)
			return
		}
	}
	logger.Error("TogglePrevTab: should never happen")
	// should never happen
}
