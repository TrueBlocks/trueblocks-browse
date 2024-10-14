package app

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
