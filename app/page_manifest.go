// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

func (a *App) ManifestPage(first, pageSize int) *types.ManifestContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.manifests.CollateAndFilter()
	first = base.Max(0, base.Min(first, len(a.manifests.Items)-1))
	last := base.Min(len(a.manifests.Items), first+pageSize)
	copy, _ := a.manifests.ShallowCopy().(*types.ManifestContainer)
	copy.Items = a.manifests.Items[first:last]
	return copy
}
