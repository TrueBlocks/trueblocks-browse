// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) NamePage(first, pageSize int) *types.NameContainer {
	// EXISTING_CODE
	namesMutex.Lock()
	defer namesMutex.Unlock()
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.names.Items)-1))
	last := base.Min(len(a.names.Items), first+pageSize)
	copy, _ := a.names.ShallowCopy().(*types.NameContainer)
	copy.Items = a.names.Items[first:last]
	return copy
}
