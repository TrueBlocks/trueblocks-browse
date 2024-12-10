// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) FetchNeighbor(first, pageSize int) *types.NeighborContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	filtered := a.neighbors.CollateAndFilter(a.GetFilter()).([]types.Transaction)
	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := a.neighbors.ShallowCopy().(*types.NeighborContainer)
	copy.Items = filtered[first:last]

	return copy
}
