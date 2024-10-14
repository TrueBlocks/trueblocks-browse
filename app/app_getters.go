package app

func (a *App) GetAppTitle() string {
	return a.session.Window.Title
}

func (a *App) IsShowing(which string) bool {
	switch which {
	case "header":
		return a.session.Toggles.Header
	case "menu":
		return a.session.Toggles.Menu
	case "help":
		return a.session.Toggles.Help
	case "footer":
		return a.session.Toggles.Footer
	}
	return false
}
