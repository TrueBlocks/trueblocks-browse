package app

import (
	"os"
	"time"
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
	if len(chain) == 0 {
		return
	}

	a.emitInfoMsg("Switching to chain", chain)

	a.session.LastChain = chain
	a.saveSession()

	a.CancelAllContexts()
	// a.project.LastUpdate = time.Time{}
	a.monitors.LastUpdate = time.Time{}
	// a.names.LastUpdate = time.Time{}
	// a.abis.LastUpdate = time.Time{}
	a.indexes.LastUpdate = time.Time{}
	a.manifests.LastUpdate = time.Time{}
	// a.status.LastUpdate = time.Time{}
	a.settings.LastUpdate = time.Time{}
	a.session.LastUpdate = time.Time{}
	a.config.LastUpdate = time.Time{}
	// a.wizard.LastUpdate = time.Time{}
	// TODO: must stop and then restart all the daemons on the new chain
	// a.daemons.LastUpdate = time.Time{}
	_ = a.Freshen()
}
