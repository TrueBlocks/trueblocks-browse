package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

func (a *App) GetFilter() *types.Filter {
	if ret, exists := a.filterMap.Load(a.GetLastRoute()); exists {
		return &ret
	}
	return &types.Filter{}
}

func (a *App) SetFilter(criteria string) {
	filter := types.Filter{
		Criteria: criteria,
	}
	a.filterMap.Store(a.GetLastRoute(), filter)
}
