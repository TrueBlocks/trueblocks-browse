package app

import (
	"os"
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

func (a *App) SetChain(chain string) {
	defer a.trackPerformance("SetChain")()
	if len(chain) == 0 || chain == a.session.LastChain {
		return
	}

	a.emitInfoMsg("Switching to chain", chain)

	a.session.LastChain = chain
	a.saveSession()

	a.CancelAllContexts()
	// a.project.LastUpdate = 0
	a.monitors.LastUpdate = 0
	// a.names.LastUpdate = 0
	// a.abis.LastUpdate = 0
	a.indexes.LastUpdate = 0
	a.manifests.LastUpdate = 0
	// a.status.LastUpdate = 0
	a.settings.LastUpdate = 0
	a.session.LastUpdate = 0
	a.config.LastUpdate = 0
	// a.wizard.LastUpdate = 0
	// TODO: must stop and then restart all the daemons on the new chain
	// a.daemons.LastUpdate = 0
	_ = a.Freshen()
}
