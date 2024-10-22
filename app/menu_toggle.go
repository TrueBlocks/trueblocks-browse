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

func (a *App) SwitchTabPrev(cd *menu.CallbackData) {
	logger.Info("SwitchTabPrev pressed")
	which := "prev"
	messages.EmitSwitchTab(a.ctx, which)
}

func (a *App) SwitchTabNext(cd *menu.CallbackData) {
	logger.Info("SwitchTabNext pressed")
	which := "next"
	messages.EmitSwitchTab(a.ctx, which)
}
