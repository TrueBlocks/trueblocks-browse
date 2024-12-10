package app

import (
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

	key := a.GetLastRoute() + "-" + a.GetLastTab()
	newState := !a.session.IsFlagOn(key)
	a.session.SetFlagOn(key, newState)
	a.saveSessionFile()

	a.emitMsg(messages.ToggleHeader, &messages.MessageMsg{
		String1: a.GetLastRoute(),
		String2: a.GetLastTab(),
		Bool:    newState,
	})
}

func (a *App) TogglePrevTab(cb *menu.CallbackData) {
	tabs := a.GetTabs()
	for i, tab := range tabs {
		if tab == a.GetLastTab() {
			prev := tabs[(i-1+len(tabs))%len(tabs)]
			a.SetLastTab(a.GetLastRoute(), prev)
			a.emitMsg(messages.Navigate, &messages.MessageMsg{String1: a.GetLastRoute(), String2: prev})
			return
		}
	}
	// should never happen
}

func (a *App) ToggleNextTab(cb *menu.CallbackData) {
	tabs := a.GetTabs()
	for i, tab := range tabs {
		if tab == a.GetLastTab() {
			next := tabs[(i+1)%len(tabs)]
			a.SetLastTab(a.GetLastRoute(), next)
			a.emitMsg(messages.Navigate, &messages.MessageMsg{String1: a.GetLastRoute(), String2: next})
			return
		}
	}
	// should never happen
}
