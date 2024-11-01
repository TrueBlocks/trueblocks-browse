package app

import "strings"

func (a *App) shouldSaveDialog() bool {
	if !a.dirty {
		return true
	}

	if response, err := a.okCancel(saveFile); err != nil {
		// there was an error, do not proceed
		a.emitErrorMsg(ErrSavingProject, err)
		return false

	} else if response.canceled {
		// user hit cancel, do not proceed
		return false

	} else if response.save {
		// saveFileDialog sends messages since it is called from FileSave, etc.
		if strings.HasPrefix(a.getFilename(), "Untitled.") {
			if saved, err := a.saveFileDialog(); err != nil || !saved {
				// there was an error or the user told us not to save.
				// in both bases, messages have been sent. do not proceed.
				return false
			}
		} else {
			if a.dirty, err = a.writeFile(a.getFullPath()); err != nil {
				a.emitErrorMsg(ErrSavingProject, err)
				return false
			}
		}
	} else {
		// do nothing - user said not to save, so just continue
	}

	return true
}
