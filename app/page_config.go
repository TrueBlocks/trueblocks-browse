// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import "github.com/TrueBlocks/trueblocks-browse/pkg/types"

// EXISTING_CODE

func (a *App) FetchConfig(first, pageSize int) *types.ConfigContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.config.CollateAndFilter()
	copy, _ := a.config.ShallowCopy().(*types.ConfigContainer)
	return copy
}
