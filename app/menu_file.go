package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
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
	a.project = types.ProjectContainer{}
	a.project.Filename = file
	a.project.Load()
	for i, history := range a.project.Items {
		if i == 0 {
			a.project.Session.LastSub = map[string]string{"/history": history.Address.Hex()}
		}
		a.HistoryPage(history.Address.Hex(), 0, 15)
	}
	a.session = a.project.Session
	messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg(a.project.Filename, "Opened"))
}

func (a *App) FileSave(cd *menu.CallbackData) {
	a.project.Filename, _ = runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           "/Users/jrush/Documents/",
		DefaultFilename:            a.project.Filename,
		Title:                      "Save File",
		CanCreateDirectories:       true,
		ShowHiddenFiles:            false,
		TreatPackagesAsDirectories: false,
		Filters: []runtime.FileFilter{
			{DisplayName: "Monitor Groups", Pattern: "*.tbx"},
		},
	})
	a.project.Session = a.session
	a.project.Save()
	messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg(a.project.Filename, "Saved"))
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	logger.Info("File SaveAs")
}
