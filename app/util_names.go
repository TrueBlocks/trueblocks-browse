package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetName(address base.Address) string {
	if name, exists := a.namesMap[address]; exists {
		return name.Name
	}
	return ""
}
