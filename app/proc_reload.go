// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// EXISTING_CODE

func (a *App) Reload() {
	defer a.trackPerformance("Reload", false)()

	route := a.session.LastRoute
	logger.InfoG("Reloading", route, "...")

	switch route {
	case "/":
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
	case "/settings":
		a.status.Updater.Reset()
		if err := a.loadStatus(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.config.Updater.Reset()
		if err := a.loadConfig(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.session.Updater.Reset()
		if err := a.loadSession(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "/daemons":
		a.daemons.Updater.Reset()
		if err := a.loadDaemons(nil, nil); err != nil {
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

	a.emitMsg(messages.Refresh, &messages.MessageMsg{
		Name:    a.daemons.FreshenController.Name,
		String1: "Refresh...",
		String2: a.daemons.FreshenController.Color,
		Num1:    1, // 1 means daemon if we need it
	})
}
