package app

import (
	"fmt"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) saveFileDialog() (bool, error) {
	if !a.dirty {
		return true, nil
	}

	fn, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           a.getFolder(),
		DefaultFilename:            a.getFile(),
		Title:                      "Save File",
		CanCreateDirectories:       true,
		ShowHiddenFiles:            false,
		TreatPackagesAsDirectories: false,
		Filters: []runtime.FileFilter{
			{DisplayName: "Projects", Pattern: "*.tbx"},
		},
	})

	if err != nil {
		a.emitErrorMsg(ErrSavingProject, err)
		return false, err

	} else if len(fn) == 0 {
		a.emitInfoMsg(a.getFullPath(), "file was not saved")
		return false, nil

	} else {
		a.CancelAllContexts()
		if _, err := a.writeFile(fn); err != nil {
			a.emitErrorMsg(ErrSavingProject, err)
			return false, err
		} else {
			a.emitInfoMsg(a.getFullPath(), "file was saved")
			return true, nil
		}
	}
}

func (a *App) writeFile(fn string) (bool, error) {
	if err := a.project.Save(fn, a.GetLastAddress()); err != nil {
		return false, fmt.Errorf("%w: %v", ErrProjectNotSaved, err)
	}
	a.dirty = false
	folder, file := filepath.Split(fn)
	a.setFolder(folder)
	a.setFile(file)
	a.saveSessionFile()
	return true, nil
}
