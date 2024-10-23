package app

import (
	"os"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
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
	a.CancelAllContexts() // cancel what's happening on the old chain
	a.globals.Chain = chain
	a.session.Chain = chain
	a.saveSession()
	a.Reload(address)
	a.abis = types.AbiContainer{}
	a.index = types.IndexContainer{}
	a.manifest = types.ManifestContainer{}
	a.monitors = types.MonitorContainer{}
	a.settings = types.SettingsGroup{}
	_ = a.Refresh()
}
