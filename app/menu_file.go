package app

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) FileNew(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	if ok := a.shouldSaveDialog(); !ok {
		return
	}

	a.newFile()
}

func (a *App) FileOpen(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	if ok := a.shouldSaveDialog(); !ok {
		return
	}

	if fn, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           a.getFolder(),
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
		a.emitErrorMsg(ErrOpeningProject, err)

	} else if len(fn) == 0 {
		a.emitInfoMsg("no file was opened", "")

	} else {
		a.CancelAllContexts()
		if _, err := a.readFile(fn); err != nil {
			a.emitErrorMsg(ErrOpeningProject, err)
		} else {
			a.emitInfoMsg(a.getFullPath(), "file was opened")
		}
	}
}

func (a *App) FileSave(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	a.dirty, _ = a.saveFileDialog()
}

func (a *App) FileSaveAs(cb *menu.CallbackData) {
	if !a.isConfigured() {
		return
	}
	a.dirty = true // force the dialog
	a.dirty, _ = a.saveFileDialog()
}
