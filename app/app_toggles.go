package app

func (a *App) IsShowing(route, tab string) bool {
	return a.session.Toggles.IsOn(route)
}

func (a *App) SetShowing(route, tab string, onOff bool) {
	a.session.Toggles.SetState(route, onOff)
	a.saveSession()
}
