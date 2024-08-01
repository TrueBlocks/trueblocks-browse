package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) FileNew(cd *menu.CallbackData) {
	logger.Info("File New")
}

func (a *App) FileOpen(cd *menu.CallbackData) {
	file, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           "/Users/jrush/Documents/",
		Title:                      "Open File",
		CanCreateDirectories:       true,
		ShowHiddenFiles:            false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
		// Filters:                    []menu.FileFilter{},
	})
	messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg(file, "Opened"))
}

func (a *App) FileSave(cd *menu.CallbackData) {
	messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg("Unknown.tbx", "Saved"))
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	logger.Info("File SaveAs")
}
