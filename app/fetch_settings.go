// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

func (a *App) FetchSettings(first, pageSize int) *types.SettingsContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	_ = a.settings.CollateAndFilter(a.filterMap)
	copy, _ := a.settings.ShallowCopy().(*types.SettingsContainer)
	return copy
}
