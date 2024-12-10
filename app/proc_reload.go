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

	logger.InfoG("Reloading", a.getLastRoute(), "...")
	switch a.getLastRoute() {

	case "", "project":
		a.project.Updater.Reset()
		if err := a.loadProject(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "history":
		a.balances.Updater.Reset()
		if err := a.loadBalances(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.incoming.Updater.Reset()
		if err := a.loadIncoming(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.outgoing.Updater.Reset()
		if err := a.loadOutgoing(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.internals.Updater.Reset()
		if err := a.loadInternals(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.charts.Updater.Reset()
		if err := a.loadCharts(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.logs.Updater.Reset()
		if err := a.loadLogs(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.statements.Updater.Reset()
		if err := a.loadStatements(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.neighbors.Updater.Reset()
		if err := a.loadNeighbors(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.traces.Updater.Reset()
		if err := a.loadTraces(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.receipts.Updater.Reset()
		if err := a.loadReceipts(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "monitors":
		a.monitors.Updater.Reset()
		if err := a.loadMonitors(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "sharing":
		a.names.Updater.Reset()
		if err := a.loadNames(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.abis.Updater.Reset()
		if err := a.loadAbis(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.pins.Updater.Reset()
		if err := a.loadPins(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.uploads.Updater.Reset()
		if err := a.loadUploads(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "unchained":
		a.indexes.Updater.Reset()
		if err := a.loadIndexes(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.manifests.Updater.Reset()
		if err := a.loadManifests(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.pins.Updater.Reset()
		if err := a.loadPins(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.uploads.Updater.Reset()
		if err := a.loadUploads(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "settings":
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
	case "daemons":
		a.daemons.Updater.Reset()
		if err := a.loadDaemons(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "wizard":
		a.wizard.Updater.Reset()
		if err := a.loadWizard(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
		// EXISTING_CODE
	default:
		address := a.getLastAddress()
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
