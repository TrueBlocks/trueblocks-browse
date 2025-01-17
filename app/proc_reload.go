// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"

// EXISTING_CODE

func (a *App) Reload() {
	defer a.trackPerformance("Reload", false)()

	route := a.session.LastRoute
	logger.InfoG("Reloading", route, "...")

	switch route {

	case "/project":
		a.project.Updater.Reset()
		if err := a.loadProject(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/monitors":
		a.monitors.Updater.Reset()
		if err := a.loadMonitors(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/names":
		a.names.Updater.Reset()
		if err := a.loadNames(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/abis":
		a.abis.Updater.Reset()
		if err := a.loadAbis(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/indexes":
		a.indexes.Updater.Reset()
		if err := a.loadIndexes(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/manifests":
		a.manifests.Updater.Reset()
		if err := a.loadManifests(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/status":
		a.status.Updater.Reset()
		if err := a.loadStatus(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/settings":
		a.settings.Updater.Reset()
		if err := a.loadSettings(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/daemons":
		a.daemons.Updater.Reset()
		if err := a.loadDaemons(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/session":
		a.session.Updater.Reset()
		if err := a.loadSession(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/config":
		a.config.Updater.Reset()
		if err := a.loadConfig(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

	case "/wizard":
		a.wizard.Updater.Reset()
		if err := a.loadWizard(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}

		// EXISTING_CODE
	default:
		address := a.GetSelected()
		// HIST-HIST
		history, _ := a.historyCache.Load(address)
		history.Updater.Reset()
		// HIST-HIST
		a.historyCache.Store(address, history)
		a.goToAddress(history.Address)
		// EXISTING_CODE
	}
}
