package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) HelpToggle(cd *menu.CallbackData) {
	logger.Info("Help Toggle")
	messages.Send(a.ctx, messages.ToggleHelp, messages.NewHelpMsg())
}
