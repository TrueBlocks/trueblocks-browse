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
	history, _ := a.historyCache.Load(address)
	// EXISTING_CODE

	_ = history.CollateAndFilter(a.filterMap)
	first = base.Max(0, base.Min(first, len(history.Items)-1))
	last := base.Min(len(history.Items), first+pageSize)
	copy, _ := history.ShallowCopy().(*types.HistoryContainer)
	copy.Balance = a.getBalance(address)
	copy.Items = history.Items[first:last]
	return copy
}
