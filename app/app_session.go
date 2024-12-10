package app

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) IsLayoutOn(layout string) bool {
	return a.session.IsFlagOn(layout)
}

func (a *App) IsHeaderOn(tab string) bool {
	key := a.GetLastRoute() + "-" + tab
	return a.session.IsFlagOn(key)
}

func (a *App) IsDaemonOn(daemon string) bool {
	return a.session.IsFlagOn(daemon)
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

// ------ access for the front end
func (a *App) SetLayoutOn(layout string, onOff bool) {
	a.session.SetFlagOn(layout, onOff)
	a.saveSessionFile()
}

func (a *App) SetHeaderOn(tab string, onOff bool) {
	key := a.GetLastRoute() + "-" + tab
	a.session.SetFlagOn(key, onOff)
	a.saveSessionFile()
}

func (a *App) SetDaemonOn(daemon string, onOff bool) {
	a.session.SetFlagOn(daemon, onOff)
	a.saveSessionFile()
}

func (a *App) SetLastRoute(route string) {
	a.session.SetRoute(route)
	a.saveSessionFile()
}

func (a *App) SetLastAddress(address string) {
	a.session.SetAddress(address)
	a.saveSessionFile()
}

func (a *App) SetLastTab(route, tab string) {
	a.session.SetTab(route, tab)
	a.saveSessionFile()
}

func (a *App) GetLastRoute() string {
	return a.session.GetRoute()
}

func (a *App) GetLastTab() string {
	return a.session.GetTab(a.GetLastRoute())
}

func (a *App) GetLastAddress() base.Address {
	return base.HexToAddress(a.session.GetAddress())
}

func (a *App) loadSessionFile() error {
	return a.session.Load()
}

func (a *App) saveSessionFile() {
	a.setWizardStr(a.getWizState().String())
	a.session.Save(a.ctx)
}
