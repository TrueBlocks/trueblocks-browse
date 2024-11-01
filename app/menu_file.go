package app

import (
	"fmt"
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
	// save current file (it will ask to save if it's dirty)
	if _, err := a.SaveFile(); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Sprintf("error saving project: %s", err.Error()),
		})
		return
	}

	a.projects = types.NewProjectContainer("Untitled.tbx", []types.HistoryContainer{})
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: "/",
	})
	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
		String1: filepath.Join(a.session.LastFolder, a.session.LastFile),
		String2: "New file created: Untitled.tbx",
	})
}

func (a *App) FileOpen(cd *menu.CallbackData) {
	// save current file (it will ask to save if it's dirty)
	if _, err := a.SaveFile(); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Sprintf("error saving project: %s", err.Error()),
		})
		return
	}

	if fn, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           a.session.LastFolder,
		DefaultFilename:            "",
		Title:                      "Open File",
		CanCreateDirectories:       true,
		ShowHiddenFiles:            false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
		Filters: []runtime.FileFilter{
			{DisplayName: "Monitor Groups", Pattern: "*.tbx"},
		},
	}); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Sprintf("error opening project: %s", err.Error()),
		})
	} else if len(fn) == 0 {
		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: "No file was opened",
		})
	} else {
		if _, err := a.LoadFile(fn); err != nil {
			messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
				String1: fmt.Sprintf("error opening project: %s", err.Error()),
			})
		} else {
			messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
				String1: filepath.Join(a.session.LastFolder, a.session.LastFile),
				String2: "Opened",
			})

		}
	}
}

func (a *App) FileSave(cd *menu.CallbackData) {
	a.SaveFile()
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	a.dirty = true
	a.SaveFile()
}

// ------------------------------------------------------------------
func (a *App) LoadFile(fn string) (bool, error) {
	newProject := types.NewProjectContainer(a.session.LastChain, []types.HistoryContainer{})
	if pF, err := newProject.Load(fn); err != nil {
		return false, fmt.Errorf("newProject::Load failed: %s", err.Error())

	} else if len(pF.Addresses) == 0 {
		logger.Info("I am here:", pF.String())
		return false, fmt.Errorf("project file contains no records: %s", fn)

	} else {
		a.CancelAllContexts()
		a.HistoryCache = &types.HistoryMap{}
		histories := []types.HistoryContainer{}
		for _, address := range pF.Addresses {
			history := types.NewHistoryContainer(a.session.LastChain, []coreTypes.Transaction{}, address)
			histories = append(histories, history)
			a.HistoryCache.Store(address, history)
		}
		a.projects = types.NewProjectContainer(a.session.LastChain, histories)

		a.session.LastFolder, a.session.LastFile = filepath.Split(fn)
		a.session.LastSub[pF.Selected.Hex()] = pF.Selected.Hex()
		a.saveSession()

		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: filepath.Join(a.session.LastFolder, a.session.LastFile),
			String2: "Opened",
		})

		var wg sync.WaitGroup
		for _, history := range a.projects.Items {
			wg.Add(1)
			go a.loadHistory(history.Address, &wg, nil)
		}
		wg.Wait()

		return true, nil
	}
}

func (a *App) WriteFile(fn string) (bool, error) {
	if err := a.projects.Save(fn, a.GetAddress()); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Sprintf("error writing project: %s", err.Error()),
		})
		return false, err
	}
	a.session.LastFolder, a.session.LastFile = filepath.Split(fn)
	a.saveSession()
	a.dirty = false
	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
		String1: filepath.Join(a.session.LastFolder, a.session.LastFile),
		String2: "Saved",
	})
	return true, nil
}

func (a *App) SaveFile() (bool, error) {
	if !a.dirty {
		return a.WriteFile(filepath.Join(a.session.LastFolder, a.session.LastFile))
	}

	var fn string
	var err error
	if isTesting {
		fn = filepath.Join(a.session.LastFolder, a.session.LastFile)
	} else {
		fn, err = runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
			DefaultDirectory:           a.session.LastFolder,
			DefaultFilename:            a.session.LastFile,
			Title:                      "Save File",
			CanCreateDirectories:       true,
			ShowHiddenFiles:            false,
			TreatPackagesAsDirectories: false,
			Filters: []runtime.FileFilter{
				{DisplayName: "Projects", Pattern: "*.tbx"},
			},
		})
	}

	if err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Sprintf("error saving project: %s", err.Error()),
		})
		return false, err
	} else if len(fn) == 0 {
		messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
			String1: "User hit escape...file not saved",
		})
		return false, nil
	} else {
		return a.WriteFile(fn)
	}
}

func (a *App) Filename() string {
	return filepath.Join(a.session.LastFolder, a.session.LastFile)
}
