// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

func (a *App) FetchIndex(first, pageSize int) *types.IndexContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	filtered := a.indexes.CollateAndFilter(a.filterMap).([]coreTypes.ChunkStats)
	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := a.indexes.ShallowCopy().(*types.IndexContainer)
	copy.Items = filtered[first:last]

	return copy
}
