package app

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
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

func (a *App) getChain() string {
	return a.session.GetChain()
}

func (a *App) getFolder() string {
	return a.session.GetFolder()
}

func (a *App) getFile() string {
	return a.session.GetFile()
}

func (a *App) getWindow() types.Window {
	return a.session.GetWindow()
}

func (a *App) wizardStr() string {
	return a.session.GetWizardStr()
}

func (a *App) cleanWindow(ctx context.Context) (types.Window, error) {
	return a.session.CleanWindowSize(ctx)
}

func (a *App) setChain(chain string) {
	a.session.SetChain(chain)
}

func (a *App) setFolder(folder string) {
	a.session.SetFolder(folder)
}

func (a *App) setFile(file string) {
	a.session.SetFile(file)
}

func (a *App) setWindow(w types.Window) {
	a.session.SetWindow(w)
}

func (a *App) setWizardStr(wizStr string) {
	a.session.SetWizardStr(wizStr)
}

func (a *App) saveSession() {
	if !isTesting {
		var w types.Window
		w.X, w.Y = runtime.WindowGetPosition(a.ctx)
		w.Width, w.Height = runtime.WindowGetSize(a.ctx)
		w.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
		a.setWindow(w)
	}
	// we serialize the wizard state in a string
	a.setWizardStr(string(a.wizard.State))
	a.session.Save()
}
