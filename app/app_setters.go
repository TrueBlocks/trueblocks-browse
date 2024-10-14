package app

import (
	"os"
)

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

func (a *App) SetEnv(key, value string) {
	os.Setenv(key, value)
}

func (a *App) SetRoute(route, subRoute string) {
	a.session.SetRoute(route, subRoute)
	a.saveSession()
}
