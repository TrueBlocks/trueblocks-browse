package app

func (a *App) Navigate(route string) {
	if route != "wizard" && !a.isConfigured() {
		route = "wizard"
	}

	a.SetLastRoute(route)
	a.emitNavigateMsg(route)
}
