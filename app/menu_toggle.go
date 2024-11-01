package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) HeaderToggle(cd *menu.CallbackData) {
	which := "header"
	messages.EmitMessage(a.ctx, messages.ToggleLayout, &messages.MessageMsg{
		String1: which,
	})
}

func (a *App) MenuToggle(cd *menu.CallbackData) {
	which := "menu"
	messages.EmitMessage(a.ctx, messages.ToggleLayout, &messages.MessageMsg{
		String1: which,
	})
}

func (a *App) HelpToggle(cd *menu.CallbackData) {
	which := "help"
	messages.EmitMessage(a.ctx, messages.ToggleLayout, &messages.MessageMsg{
		String1: which,
	})
}

func (a *App) FooterToggle(cd *menu.CallbackData) {
	which := "footer"
	messages.EmitMessage(a.ctx, messages.ToggleLayout, &messages.MessageMsg{
		String1: which,
	})
}

func (a *App) AccordionToggle(cd *menu.CallbackData) {
	route := a.GetRoute()
	route = strings.TrimPrefix(route, "/")
	parts := strings.Split(route, "/")
	route = parts[0]
	if route == "" {
		route = "project"
	}
	messages.EmitMessage(a.ctx, messages.ToggleHeader, &messages.MessageMsg{
		String1: route,
	})
}

func (a *App) SwitchTabPrev(cd *menu.CallbackData) {
	which := "prev"
	messages.EmitMessage(a.ctx, messages.SwitchTab, &messages.MessageMsg{String1: which})
}

func (a *App) SwitchTabNext(cd *menu.CallbackData) {
	which := "next"
	messages.EmitMessage(a.ctx, messages.SwitchTab, &messages.MessageMsg{String1: which})
}
