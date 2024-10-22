package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) HeaderToggle(cd *menu.CallbackData) {
	logger.Info("HeaderToggle")
	which := "header"
	messages.EmitToggle(a.ctx, which, "")
}

func (a *App) MenuToggle(cd *menu.CallbackData) {
	logger.Info("MenuToggle")
	which := "menu"
	messages.EmitToggle(a.ctx, which, "")
}

func (a *App) HelpToggle(cd *menu.CallbackData) {
	logger.Info("HelpToggle")
	which := "help"
	messages.EmitToggle(a.ctx, which, "")
}

func (a *App) FooterToggle(cd *menu.CallbackData) {
	logger.Info("FooterToggle")
	which := "footer"
	messages.EmitToggle(a.ctx, which, "")
}

func (a *App) AccordionToggle(cd *menu.CallbackData) {
	logger.Info("AccordionToggle")
	route := a.GetRoute()
	route = strings.TrimPrefix(route, "/")
	parts := strings.Split(route, "/")
	route = parts[0]
	if route == "" {
		route = "project"
	}
	messages.EmitToggle(a.ctx, "", route)
}

func (a *App) SwithTabHome(cd *menu.CallbackData) {
	logger.Info("SwithTabHome pressed")
	which := "home"
	messages.EmitSwitchTab(a.ctx, which)
}

func (a *App) SwithTabPrev(cd *menu.CallbackData) {
	logger.Info("SwithTabPrev pressed")
	which := "previous"
	messages.EmitSwitchTab(a.ctx, which)
}

func (a *App) SwithTabNext(cd *menu.CallbackData) {
	logger.Info("SwithTabNext pressed")
	which := "next"
	messages.EmitSwitchTab(a.ctx, which)
}

func (a *App) SwithTabEnd(cd *menu.CallbackData) {
	logger.Info("SwithTabEnd pressed")
	which := "end"
	messages.EmitSwitchTab(a.ctx, which)
}
