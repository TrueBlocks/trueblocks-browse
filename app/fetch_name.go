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

func (a *App) FetchName(first, pageSize int) *types.NameContainer {
	// EXISTING_CODE
	namesMutex.Lock()
	defer namesMutex.Unlock()
	// EXISTING_CODE

	filtered := a.names.CollateAndFilter(a.filterMap).([]coreTypes.Name)
	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := a.names.ShallowCopy().(*types.NameContainer)
	copy.Items = filtered[first:last]

	return copy
}
