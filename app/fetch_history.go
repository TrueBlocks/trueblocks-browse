// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) FetchHistory(first, pageSize int) *types.HistoryContainer {
	// EXISTING_CODE
	address := a.GetSelected()
	// HIST-HIST
	history, _ := a.historyCache.Load(address)
	// EXISTING_CODE

	filtered := history.CollateAndFilter(a.filterMap).([]types.Transaction)
	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := history.ShallowCopy().(*types.HistoryContainer)
	copy.Items = filtered[first:last]

	return copy
}
