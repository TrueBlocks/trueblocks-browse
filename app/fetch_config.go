// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
)

// EXISTING_CODE

func (a *App) FetchConfig(first, pageSize int) *types.ConfigContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	filtered := a.config.CollateAndFilter(a.filterMap).([]configTypes.ChainGroup)
	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := a.config.ShallowCopy().(*types.ConfigContainer)
	copy.Items = filtered[first:last]

	return copy
}
