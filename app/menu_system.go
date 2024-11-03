package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) SystemAbout(cb *menu.CallbackData) {
	logger.Info("This is the about box plus some other stuff")
}

func (a *App) SystemQuit(cb *menu.CallbackData) {
	a.dirty, _ = a.saveFileDialog()
	runtime.Quit(a.ctx)
}
