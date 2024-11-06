// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) IndexPage(first, pageSize int) *types.IndexContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.indexes.Items)-1))
	last := base.Min(len(a.indexes.Items), first+pageSize)
	copy, _ := a.indexes.ShallowCopy().(*types.IndexContainer)
	copy.Items = a.indexes.Items[first:last]
	return copy
}
