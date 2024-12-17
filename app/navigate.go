package app

func (a *App) Navigate(route string) {
	if route != "wizard" && !a.isConfigured() {
		route = "wizard"
	}

	a.setLastRoute(route)
	a.emitNavigateMsg(route, a.getLastTab(route))
}
