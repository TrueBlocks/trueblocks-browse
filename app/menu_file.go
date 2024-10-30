package app

import (
	"path/filepath"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) FileNew(cd *menu.CallbackData) {
	a.projects = types.NewProjectContainer("Untitled.tbx", &types.HistoryMap{})
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: "/",
	})
	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
		String1: a.projects.Filename,
		String2: "Created",
	})
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
		save := a.FreshenController.Sleep
		defer func() { a.FreshenController.Sleep = save }()
		a.FreshenController.Sleep = 1000
		a.sessions.LastFile = file
		a.saveSession()

		a.CancelAllContexts()
		a.projects = types.NewProjectContainer(file, &types.HistoryMap{})
		newProject := types.ProjectContainer{
			Filename: file,
		}
		if err := newProject.Load(); err != nil {
			logger.Error("Error loading project", err.Error(), a.sessions.Session.LastFile)
			newProject.Session = a.sessions.Session // backup plan
			newProject.Session.Wizard.State = coreTypes.Error
		}
		copyFields := func(dest, src *coreTypes.Session) {
			dest.LastChain = src.LastChain
			dest.LastFile = src.LastFile
			dest.LastFolder = src.LastFolder
			dest.LastRoute = src.LastRoute
			dest.LastSub = src.LastSub
			dest.Toggles = src.Toggles
			dest.Window = src.Window
			dest.Wizard = src.Wizard
		}
		copyFields(&a.sessions.Session, &newProject.Session)

		a.sessions = types.NewSessionContainer(a.Chain, &a.sessions.Session)
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
			String1: a.projects.Filename,
			String2: "Opened",
		})
	} else {
		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: a.projects.Filename,
			String2: "Not opened",
		})
	}
}

func (a *App) FileSave(cd *menu.CallbackData) {
	a.saveSession()
	fn, _ := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           a.sessions.Session.LastFolder,
		DefaultFilename:            a.projects.Filename,
		Title:                      "Save File",
		CanCreateDirectories:       true,
		ShowHiddenFiles:            false,
		TreatPackagesAsDirectories: false,
		Filters: []runtime.FileFilter{
			{DisplayName: "Projects", Pattern: "*.tbx"},
		},
	})
	if len(fn) > 0 {
		a.sessions.Session.LastFolder, a.sessions.Session.LastFile = filepath.Split(fn)
		a.sessions.Session.LastRoute = "/"

		a.projects.Session = a.sessions.Session
		a.projects.Filename = fn
		a.projects.Save()
		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: a.projects.Filename,
			String2: "Saved",
		})
	} else {
	}
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	a.saveSession()
}
