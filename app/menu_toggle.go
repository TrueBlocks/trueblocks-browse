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

	key := a.GetRawRoute() + "-" + a.GetLastTab()
	newState := !a.session.IsFlagOn(key)
	a.session.SetFlagOn(key, newState)
	a.saveSession()

	a.emitMsg(messages.ToggleHeader, &messages.MessageMsg{
		String1: a.GetRawRoute(),
		String2: a.GetLastTab(),
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
