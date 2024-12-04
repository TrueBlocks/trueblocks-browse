package app

import (
	"os"
)

func (a *App) SetEnv(key, value string) {
	os.Setenv(key, value)
}

func (a *App) SetChain(newChain string) {
	defer a.trackPerformance("SetChain", false)()

	oldChain := a.session.LastChain
	if len(newChain) == 0 || newChain == oldChain {
		return
	}

	a.emitInfoMsg("Switching to chain", newChain)

	a.session.LastChain = newChain
	a.saveSession()

	a.CancelAllContexts()
	a.project.Updater.SetChain(oldChain, newChain)
	a.monitors.Updater.SetChain(oldChain, newChain)
	// a.names.Updater.SetChain(oldChain, newChain)
	// a.abis.Updater.SetChain(oldChain, newChain)
	a.indexes.Updater.SetChain(oldChain, newChain)
	a.manifests.Updater.SetChain(oldChain, newChain)
	a.status.Updater.SetChain(oldChain, newChain)
	a.session.Updater.SetChain(oldChain, newChain)
	a.config.Updater.SetChain(oldChain, newChain)
	a.wizard.Updater.SetChain(oldChain, newChain)
	// TODO: must stop and then restart all the daemons on the new newChain
	a.daemons.Updater.SetChain(oldChain, newChain)
	_ = a.Freshen()
}
