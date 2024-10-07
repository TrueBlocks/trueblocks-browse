package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) Reload(address base.Address) {
	a.ModifyProject(&ModifyData{
		Operation: "reload",
		Address:   address,
	})
	a.HistoryPage(address.Hex(), 0, 15)
	a.Refresh()
	a.loadProject(nil, nil)
}

func (a *App) ModifyProject(modData *ModifyData) {
	a.CancelContext(modData.Address)
	a.project.HistoryMap.Delete(modData.Address)
	for i, item := range a.project.Items {
		if item.Address == modData.Address {
			a.project.Items = append(a.project.Items[:i], a.project.Items[i+1:]...)
			break
		}
	}
}
