package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

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

func (a *App) GetAppTitle() string {
	return a.session.Window.Title
}

func (a *App) GetRouteAndSub() (string, string) {
	if !a.isConfigured() {
		return "/wizard"
	}

	route := a.session.LastRoute
	sub, _ := a.session.LastSub.Load(route)
	if len(sub) > 0 {
		route += "/" + sub
	}

	return route
}

func (a *App) GetLastAddress() base.Address {
	addr, _ := a.session.LastSub.Load("/history")
	return base.HexToAddress(addr)
}
