package app

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type Filter struct {
	Criteria string `json:"criteria"`
}

func (f *Filter) String() string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}

func (a *App) GetFilter(route string) Filter {
	if ret, exists := a.filterMap.Load(route); exists {
		return ret.(Filter)
	}
	return Filter{}
}

func (a *App) SetFilter(route, criteria string) {
	filter := Filter{
		Criteria: criteria,
	}
	logger.InfoBM(filter)
	a.filterMap.Store(route, filter)
}
