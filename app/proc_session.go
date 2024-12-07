package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) saveSession() {
	if !isTesting {
		var w types.Window
		w.X, w.Y = runtime.WindowGetPosition(a.ctx)
		w.Width, w.Height = runtime.WindowGetSize(a.ctx)
		w.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
		a.session.SetWindow(w)
	}
	// we serialize the wizard state in a session string
	a.session.SetWizardStr(string(a.wizard.State))
	_ = a.session.Save()
}
