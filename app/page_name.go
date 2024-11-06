// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) NamePage(first, pageSize int) *types.NameContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.names.Items)-1))
	last := base.Min(len(a.names.Items), first+pageSize)
	copy, _ := a.names.ShallowCopy().(*types.NameContainer)
	copy.Items = a.names.Items[first:last]
	return copy
}
