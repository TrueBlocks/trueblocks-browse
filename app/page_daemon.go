// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import "github.com/TrueBlocks/trueblocks-browse/pkg/types"

// EXISTING_CODE

func (a *App) FetchDaemon(first, pageSize int) *types.DaemonContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.daemons.CollateAndFilter()
	copy, _ := a.daemons.ShallowCopy().(*types.DaemonContainer)
	return copy
}
