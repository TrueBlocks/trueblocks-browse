// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

func (a *App) WizardPage(first, pageSize int) *types.WizardContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.wizard.Summarize()
	first = base.Max(0, base.Min(first, len(a.wizard.Items)-1))
	last := base.Min(len(a.wizard.Items), first+pageSize)
	copy, _ := a.wizard.ShallowCopy().(*types.WizardContainer)
	copy.Items = a.wizard.Items[first:last]
	return copy
}
