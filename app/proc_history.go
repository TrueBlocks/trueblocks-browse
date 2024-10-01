package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) Reload(addr base.Address) {
	a.CancelContexts()
	a.removeAddress(addr)
	a.HistoryPage(addr.String(), 0, 15)
	a.Refresh(false)
	a.loadProject(nil, nil)
}

func (a *App) removeAddress(addr base.Address) {
	historyMutex.Lock()
	delete(a.historyMap, addr)
	historyMutex.Unlock()
	for i, item := range a.project.Items {
		if item.Address == addr {
			a.project.Items = append(a.project.Items[:i], a.project.Items[i+1:]...)
			// a.project.MyCount--
			break
		}
	}
	// for i, item := range a.monitors.Items {
	// 	if item.Address == addr {
	// 		a.monitors.Items = append(a.monitors.Items[:i], a.monitors.Items[i+1:]...)
	// 		// a.monitors.NItems--
	// 		break
	// 	}
	// }
}
