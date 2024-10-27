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
	logger.Info("File New")
	a.project = types.NewProjectContainer("Untitled.tbx", &types.HistoryMap{}, &sync.Map{}, &sync.Map{})
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: "/",
	})
	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
		String1: a.project.Filename,
		String2: "Created",
	})
}

func (a *App) FileOpen(cd *menu.CallbackData) {
	logger.Info("File Open")
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
		save := a.FreshenController.Sleep
		defer func() { a.FreshenController.Sleep = save }()
		a.FreshenController.Sleep = 1000
		a.session.LastFile = file
		a.saveSession()

		a.CancelAllContexts()
		a.project = types.NewProjectContainer(file, &types.HistoryMap{}, &sync.Map{}, &sync.Map{})
		newProject := types.ProjectContainer{
			Filename: file,
		}
		newProject.Load()
		a.session = newProject.Session
		var wg sync.WaitGroup
		for _, history := range newProject.Items {
			wg.Add(1)
			go a.loadHistory(history.Address, &wg, nil)
		}
		wg.Wait()

		messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
			String1: "/",
		})
		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: a.project.Filename,
			String2: "Opened",
		})
	} else {
		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: a.project.Filename,
			String2: "Not opened",
		})
	}
}

func (a *App) FileSave(cd *menu.CallbackData) {
	logger.Info("File Save")
	a.saveSession()
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
	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
		String1: a.project.Filename,
		String2: "Saved",
	})
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	logger.Info("File SaveAs")
	a.saveSession()
}
