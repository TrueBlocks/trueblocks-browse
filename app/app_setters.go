package app

import (
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) SetShowing(which string, onOff bool) {
	a.session.Toggles.SetState(which, onOff)
	a.saveSession()
}

func (a *App) SetEnv(key, value string) {
	os.Setenv(key, value)
}

func (a *App) SetRoute(route, subRoute string) {
	a.session.SetRoute(route, subRoute)
	a.saveSession()
}

func (a *App) SetChain(chain string, address base.Address) {
	a.emitInfoMsg("Switching to chain", chain)
	a.CancelAllContexts()
	a.Chain = chain
	a.session.LastChain = chain
	a.saveSession()
	a.GoToAddress(address)
	a.monitors = types.MonitorContainer{}
	a.names = types.NameContainer{}
	a.abis = types.AbiContainer{}
	a.indexes = types.IndexContainer{}
	a.manifests = types.ManifestContainer{}
	a.status = types.StatusContainer{}
	a.settings = types.SettingsGroup{}
	_ = a.Refresh()
}
