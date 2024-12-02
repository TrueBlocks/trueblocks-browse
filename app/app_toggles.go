package app

func (a *App) IsShowing(route string) bool {
	return a.session.Toggles.IsOn(route)
}

func (a *App) SetShowing(route string, onOff bool) {
	a.session.Toggles.SetState(route, onOff)
	a.saveSession()
}

