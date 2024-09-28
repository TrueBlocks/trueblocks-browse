package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) Reload(addr base.Address) {
	a.CancleContexts()
	historyMutex.Lock()
	delete(a.historyMap, addr)
	historyMutex.Unlock()
	a.HistoryPage(addr.String(), 0, 15)
	a.removeAddress(addr)
	a.Refresh(false)
}

func (a *App) removeAddress(addr base.Address) {
	for i, item := range a.portfolio.Items {
		if item.Address == addr {
			a.portfolio.Items = append(a.portfolio.Items[:i], a.portfolio.Items[i+1:]...)
			// a.portfolio.MyCount--
			break
		}
	}
	for i, item := range a.monitors.Items {
		if item.Address == addr {
			a.monitors.Items = append(a.monitors.Items[:i], a.monitors.Items[i+1:]...)
			// a.monitors.NItems--
			break
		}
	}
}
