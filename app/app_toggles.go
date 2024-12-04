package app

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"

func (a *App) IsLayoutOn(layout string) bool {
	return a.session.Toggles.IsOn(layout)
}

func (a *App) IsDaemonOn(daemon string) bool {
	return a.session.Toggles.IsOn(daemon)
}

func (a *App) IsHeaderOn(header, tab string) bool {
	return a.session.Toggles.IsOn(header + "-" + tab)
}

func (a *App) SetLayoutOn(layout string, onOff bool) {
	a.session.Toggles.SetState(layout, onOff)
	a.saveSession()
}

func (a *App) SetDaemonOn(daemon string, onOff bool) {
	a.session.Toggles.SetState(daemon, onOff)
	a.saveSession()
}

func (a *App) SetHeaderOn(route, tab string, onOff bool) {
	a.session.Toggles.SetState(route+"-"+tab, onOff)
	a.saveSession()
}

func (a *App) TabSwitched(route, tab string) {
	logger.InfoBB("TabSwitched", route, tab)
	a.session.SetTab(route, tab)
}

func (a *App) GetActiveTab(route string) string {
	ret, _ := a.session.LastTab.Load(route)
	return ret
}
