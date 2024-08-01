package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) HelpToggle(cd *menu.CallbackData) {
	logger.Info("Help Toggle")
	runtime.EventsEmit(a.ctx, "helpToggle")
}
