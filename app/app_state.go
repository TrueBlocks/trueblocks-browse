package app

func (a *App) GetLast(which string) string {
	switch which {
	case "route":
		return a.GetSession().LastRoute
	case "tab":
		return a.GetSession().LastTab
	case "address":
		return a.GetSession().LastAddress
	}
	return "Unknown"
}

func (a *App) SetLast(which, value string) {
	switch which {
	case "route":
		a.GetSession().LastRoute = value
	case "tab":
		a.GetSession().LastTab = value
	case "address":
		a.GetSession().LastAddress = value
	}
	a.GetSession().Save()
}