package app

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"

func (a *App) Reload() {
	defer a.trackPerformance("Reload", false)()

	route := a.session.LastRoute
	logger.InfoG("Reloading", route, "...")

	switch route {
	case "/names":
		a.names.Updater.Reset()
		if err := a.loadNames(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	case "/monitors":
		a.monitors.Updater.Reset()
		if err := a.loadMonitors(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	default:
		address := a.GetSelected()
		// HIST-HIST
		history, _ := a.historyCache.Load(address)
		history.Updater.Reset()
		// HIST-HIST
		a.historyCache.Store(address, history)
		a.goToAddress(history.Address)
	}
}
