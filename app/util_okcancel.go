package app

import "github.com/wailsapp/wails/v2/pkg/runtime"

type OkCancel struct {
	title   string
	msg     string
	options []string
	// output
	save     bool
	canceled bool
}

var saveFile = OkCancel{
	title:   "Unsaved Changes",
	msg:     "File has changed. Save it?",
	options: []string{"Save", "Don't Save", "Cancel"},
}

func (a *App) okCancel(in OkCancel) (OkCancel, error) {
	if !a.dirty {
		return OkCancel{}, nil
	}

	a.CancelAllContexts()
	if response, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   in.title,
		Message: in.msg,
		Buttons: in.options,
	}); err != nil {
		return OkCancel{}, err
	} else if response == "Don't Save" {
		return OkCancel{}, nil
	} else if response == "Cancel" {
		return OkCancel{canceled: true}, nil
	} else {
		return OkCancel{save: true}, nil
	}
}
