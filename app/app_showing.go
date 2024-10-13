package app

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

func (a *App) SetShowing(which string, onOff bool) {
	switch which {
	case "header":
		a.session.Toggles.Header = onOff
	case "menu":
		a.session.Toggles.Menu = onOff
	case "help":
		a.session.Toggles.Help = onOff
	case "footer":
		a.session.Toggles.Footer = onOff
	}
	a.saveSession()
}
