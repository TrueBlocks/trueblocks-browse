// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) HistoryPage(first, pageSize int) *types.HistoryContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	address := a.GetSelected()
	_, exists := a.historyCache.Load(address)
	if !exists {
		return &types.HistoryContainer{}
	}

	first = base.Max(0, base.Min(first, a.txCount(address)-1))
	last := base.Min(a.txCount(address), first+pageSize)
	history, _ := a.historyCache.Load(address)
	history.Summarize()
	copy := history.ShallowCopy().(*types.HistoryContainer)
	copy.Balance = a.getBalance(address)
	copy.Items = history.Items[first:last]
	return copy
}
