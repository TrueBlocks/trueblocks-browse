package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) SystemAbout(cd *menu.CallbackData) {
	logger.Info("This is the about box plus some other stuff")
}

func (a *App) SystemQuit(cd *menu.CallbackData) {
	runtime.Quit(a.ctx)
}
