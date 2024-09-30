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
		Filters: []runtime.FileFilter{
			{DisplayName: "Monitor Groups", Pattern: "*.tbx"},
		},
	})
	a.portfolio.Filename = file
	a.portfolio.Load()
	messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg(a.portfolio.Filename, "Opened"))
}

func (a *App) FileSave(cd *menu.CallbackData) {
	a.portfolio.Filename, _ = runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           "/Users/jrush/Documents/",
		DefaultFilename:            a.portfolio.Filename,
		Title:                      "Save File",
		CanCreateDirectories:       true,
		ShowHiddenFiles:            false,
		TreatPackagesAsDirectories: false,
		Filters: []runtime.FileFilter{
			{DisplayName: "Monitor Groups", Pattern: "*.tbx"},
		},
	})
	a.portfolio.Save()
	messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg(a.portfolio.Filename, "Saved"))
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	logger.Info("File SaveAs")
}
