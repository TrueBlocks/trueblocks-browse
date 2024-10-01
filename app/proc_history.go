package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) Reload(address base.Address) {
	a.CancelContexts()
	a.removeAddress(address)
	a.HistoryPage(address.Hex(), 0, 15)
	a.Refresh(false)
	a.loadProject(nil, nil)
}

func (a *App) removeAddress(address base.Address) {
	a.closeFile(address)
	for i, item := range a.project.Items {
		if item.Address == address {
			a.project.Items = append(a.project.Items[:i], a.project.Items[i+1:]...)
			// a.project.NOpenFiles--
			break
		}
	}
}
