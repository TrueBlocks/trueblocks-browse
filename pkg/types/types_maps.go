package types

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/maps"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type BoolMap = maps.Map[string, bool]
type StringMap = maps.Map[string, string]
type HistoryMap = maps.Map[base.Address, HistoryContainer]
type FilterMap = maps.Map[string, Filter]
