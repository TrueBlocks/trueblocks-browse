package app

func (a *App) IsShowing(route, tab string) bool {
	return a.session.Toggles.IsOn(route)
}

func (a *App) SetShowing(route, tab string, onOff bool) {
	a.session.Toggles.SetState(route, onOff)
	a.saveSession()
}

func (a *App) TabSwitched(route, tab string) {
	logger.InfoBW("TabSwitched", route, tab)
	// a.session.Toggles.SetState(route, true)
	// a.saveSession()
}

func (a *App) GetActiveTab(route string) string {
	return ""
	// return a.session.Toggles.GetActiveTab(route)
}

