package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) IsLayoutOn(layout string) bool {
	return a.session.IsFlagOn(layout)
}

func (a *App) SetLayoutOn(layout string, onOff bool) {
	a.session.SetFlagOn(layout, onOff)
	a.saveSession()
}

func (a *App) IsHeaderOn(route, tab string) bool {
	key := route + "-" + tab
	return a.session.IsFlagOn(key)
}

func (a *App) SetHeaderOn(route, tab string, onOff bool) {
	key := route + "-" + tab
	a.session.SetFlagOn(key, onOff)
	a.saveSession()
}

func (a *App) IsDaemonOn(daemon string) bool {
	return a.session.IsFlagOn(daemon)
}

func (a *App) SetDaemonOn(daemon string, onOff bool) {
	a.session.SetFlagOn(daemon, onOff)
	a.saveSession()
}

func (a *App) SetLastRoute(route, address string) {
	a.session.SetRoute(route)
	a.session.SetAddress(address)
	a.saveSession()
}

func (a *App) GetRawRoute() string {
	return strings.Trim(a.session.GetRoute(), "/")
}

func (a *App) GetLastRoute() string {
	if !a.isConfigured() {
		return "/wizard"
	}

	route, addr := a.session.GetRoute(), a.session.GetAddress()
	if len(addr) > 0 {
		route += "/" + addr
	}

	return route
}

func (a *App) GetLastAddress() base.Address {
	return base.HexToAddress(a.session.GetAddress())
}

func (a *App) SetLastTab(route, tab string) {
	a.session.SetTab(route, tab)
}

func (a *App) GetLastTab() string {
	route := a.GetRawRoute()
	return a.session.GetTab(route)
}
