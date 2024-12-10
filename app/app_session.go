package app

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// --------------------------------------------
func (a *App) GetLastRoute() string {
	return a.getLastRoute()
}

func (a *App) SetLastRoute(route string) {
	a.setLastRoute(route)
}

// --------------------------------------------
func (a *App) GetLastTab(route string) string {
	return a.getLastTab(route)
}

func (a *App) SetLastTab(route, tab string) {
	a.setLastTab(route, tab)
}

// --------------------------------------------
func (a *App) GetLastAddress() base.Address {
	return a.getLastAddress()
}

func (a *App) SetLastAddress(address string) {
	a.setLastAddress(address)
}

// --------------------------------------------
func (a *App) IsLayoutOn(layout string) bool {
	return a.session.IsFlagOn(layout)
}

func (a *App) SetLayoutOn(layout string, onOff bool) {
	a.session.SetFlagOn(layout, onOff)
	a.saveSessionFile()
}

// --------------------------------------------
func (a *App) IsHeaderOn(route, tab string) bool {
	key := route + "-" + tab
	return a.session.IsFlagOn(key)
}

func (a *App) SetHeaderOn(route, tab string, onOff bool) {
	key := route + "-" + tab
	a.session.SetFlagOn(key, onOff)
	a.saveSessionFile()
}

// --------------------------------------------
func (a *App) IsDaemonOn(daemon string) bool {
	return a.session.IsFlagOn(daemon)
}

func (a *App) SetDaemonOn(daemon string, onOff bool) {
	a.session.SetFlagOn(daemon, onOff)
	a.saveSessionFile()
}

// not exposed to the front end
// --------------------------------------------
func (a *App) getLastRoute() string {
	return a.session.GetRoute()
}

func (a *App) setLastRoute(route string) {
	a.session.SetRoute(route)
	a.saveSessionFile()
}

// --------------------------------------------
func (a *App) getLastTab(route string) string {
	ret := a.session.GetTab(route)
	if ret == "" {
		ret = a.GetTabs()[0]
	}
	return ret
}

func (a *App) setLastTab(route, tab string) {
	a.session.SetTab(route, tab)
	a.saveSessionFile()
}

// --------------------------------------------
func (a *App) getLastAddress() base.Address {
	return base.HexToAddress(a.session.GetAddress())
}

func (a *App) setLastAddress(address string) {
	a.session.SetAddress(address)
	a.saveSessionFile()
}

// --------------------------------------------
func (a *App) getChain() string {
	return a.session.GetChain()
}

func (a *App) setChain(chain string) {
	a.session.SetChain(chain)
}

// --------------------------------------------
func (a *App) getFolder() string {
	return a.session.GetFolder()
}

func (a *App) setFolder(folder string) {
	a.session.SetFolder(folder)
}

// --------------------------------------------
func (a *App) getFile() string {
	return a.session.GetFile()
}

func (a *App) setFile(file string) {
	a.session.SetFile(file)
}

// --------------------------------------------
func (a *App) getWindow() types.Window {
	return a.session.GetWindow()
}

func (a *App) setWindow(w types.Window) {
	a.session.SetWindow(w)
}

// --------------------------------------------
func (a *App) wizardStr() string {
	return a.session.GetWizardStr()
}

func (a *App) setWizardStr(wizStr string) {
	a.session.SetWizardStr(wizStr)
}

// --------------------------------------------
func (a *App) cleanWindow(ctx context.Context) (types.Window, error) {
	return a.session.CleanWindowSize(ctx)
}

// --------------------------------------------
func (a *App) loadSessionFile() error {
	return a.session.Load(a.ctx)
}

func (a *App) saveSessionFile() {
	a.setWizardStr(a.getWizState().String())
	a.session.Save(a.ctx)
}
