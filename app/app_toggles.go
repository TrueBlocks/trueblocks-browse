package app

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"

func (a *App) IsShowing(route, tab string) bool {
	return a.session.Toggles.IsOn(route)
}

func (a *App) IsShowing2(route string) bool {
	return a.session.Toggles.IsOn(route)
}

func (a *App) SetShowing(route, tab string, onOff bool) {
	a.session.Toggles.SetState(route, onOff)
	a.saveSession()
}

func (a *App) SetShowing2(route string, onOff bool) {
	a.session.Toggles.SetState(route, onOff)
	a.saveSession()
}

func (a *App) TabSwitched(route, tab string) {
	logger.InfoBB("TabSwitched", route, tab)
	a.session.SetTab(route, tab)
}

func (a *App) GetActiveTab(route string) string {
	if a.session.LastTab != nil {
		return a.session.LastTab[route]
	}
	return ""
}
