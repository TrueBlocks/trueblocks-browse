// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) FetchReceipt(first, pageSize int) *types.ReceiptContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	filtered := a.receipts.CollateAndFilter(a.GetFilter()).([]types.Transaction)
	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := a.receipts.ShallowCopy().(*types.ReceiptContainer)
	copy.Items = filtered[first:last]

	return copy
}
