package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) GetFilter(route string) types.Filter {
	if ret, exists := a.filterMap.Load(route); exists {
		return ret
	}
	return types.Filter{}
}

func (a *App) SetFilter(route, criteria string) {
	filter := types.Filter{
		Criteria: criteria,
	}
	logger.InfoBM(route, filter)
	a.filterMap.Store(route, filter)
}
