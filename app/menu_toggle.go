package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) HeaderToggle(cd *menu.CallbackData) {
	messages.EmitToggle(a.ctx, "header")
}

func (a *App) MenuToggle(cd *menu.CallbackData) {
	messages.EmitToggle(a.ctx, "menu")
}

func (a *App) HelpToggle(cd *menu.CallbackData) {
	messages.EmitToggle(a.ctx, "help")
}

func (a *App) FooterToggle(cd *menu.CallbackData) {
	messages.EmitToggle(a.ctx, "footer")
}
