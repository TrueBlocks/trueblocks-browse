package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) HeaderToggle(cb *menu.CallbackData) {
	which := "header"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) MenuToggle(cb *menu.CallbackData) {
	which := "menu"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) HelpToggle(cb *menu.CallbackData) {
	which := "help"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) FooterToggle(cb *menu.CallbackData) {
	which := "footer"
	a.emitMsg(messages.ToggleLayout, &messages.MessageMsg{String1: which})
}

func (a *App) AccordionToggle(cb *menu.CallbackData) {
	route := a.GetRoute()
	route = strings.TrimPrefix(route, "/")
	parts := strings.Split(route, "/")
	route = parts[0]
	if route == "" {
		route = "project"
	}
	a.emitMsg(messages.ToggleAccordion, &messages.MessageMsg{String1: route})
}

func (a *App) SwitchTabPrev(cb *menu.CallbackData) {
	which := "prev"
	a.emitMsg(messages.SwitchTab, &messages.MessageMsg{String1: which})
}

func (a *App) SwitchTabNext(cb *menu.CallbackData) {
	which := "next"
	a.emitMsg(messages.SwitchTab, &messages.MessageMsg{String1: which})
}
