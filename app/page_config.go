// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import "github.com/TrueBlocks/trueblocks-browse/pkg/types"

// EXISTING_CODE

func (a *App) ConfigPage(first, pageSize int) *types.ConfigContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.config.Summarize()
	copy, _ := a.config.ShallowCopy().(*types.ConfigContainer)
	return copy
}
