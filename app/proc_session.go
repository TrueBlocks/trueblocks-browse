package app

import "github.com/wailsapp/wails/v2/pkg/runtime"

func (a *App) saveSession() {
	if !isTesting {
		a.session.Window.X, a.session.Window.Y = runtime.WindowGetPosition(a.ctx)
		a.session.Window.Width, a.session.Window.Height = runtime.WindowGetSize(a.ctx)
		a.session.Window.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	}
	// we serialize the wizard state in a session string
	a.session.WizardStr = string(a.wizard.State)
	_ = a.session.Save()
}
