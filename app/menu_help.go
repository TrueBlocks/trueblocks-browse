package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) HelpToggle(cd *menu.CallbackData) {
	debugMsg("Help Toggle")
	messages.EmitHelp(a.ctx)
}
