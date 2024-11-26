package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
)

func (a *App) ModifyProject(modData *ModifyData) {
	switch crud.OpFromString(modData.Operation) {
	case crud.Delete:
		a.cancelContext(modData.Address)
		// HIST-PROJ
		a.historyCache.Delete(modData.Address)
		a.dirty = true
		a.emitInfoMsg(a.getFullPath(), fmt.Sprint("deleted address", modData.Address.Hex()))
	}
}
