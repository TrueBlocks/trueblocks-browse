package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) FileNew(cd *menu.CallbackData) {
	a.project = types.NewProjectContainer("Untitled.tbx", &types.HistoryMap{}, &sync.Map{}, &sync.Map{})
	messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/"))
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

	if len(file) > 0 {
		newProject := types.ProjectContainer{
			Filename: file,
		}
		newProject.Load()
		a.session = newProject.Session
		for _, history := range newProject.Items {
			history.Balance = a.getBalance(history.Address)
			if loaded, ok := a.project.HistoryMap.Load(history.Address); ok && loaded.NItems == loaded.NTotal {
				history.Items = loaded.Items
				newProject.HistoryMap.Store(history.Address, history)
				messages.Send(a.ctx,
					messages.Completed,
					messages.NewProgressMsg(int64(a.txCount(history.Address)), int64(a.txCount(history.Address)), history.Address),
				)
			} else {
				go a.HistoryPage(history.Address.Hex(), -1, 15)
			}
		}

		messages.Send(a.ctx, messages.Navigate, messages.NewNavigateMsg("/"))
		messages.Send(a.ctx, messages.Document, messages.NewDocumentMsg(a.project.Filename, "Opened"))
	}
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
