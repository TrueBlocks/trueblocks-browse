package app

func (a *App) IsShowing(which string) bool {
	return a.session.Toggles.IsOn(which)
}

func (a *App) SetShowing(which string, onOff bool) {
	a.session.Toggles.SetState(which, onOff)
	a.saveSession()
}
