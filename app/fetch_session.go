// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

func (a *App) FetchSession(first, pageSize int) *types.SessionContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.session.CollateAndFilter()
	copy, _ := a.session.ShallowCopy().(*types.SessionContainer)
	return copy
}
