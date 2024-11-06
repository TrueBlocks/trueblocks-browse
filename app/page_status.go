// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) StatusPage(first, pageSize int) *types.StatusContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.status.Items)-1))
	last := base.Min(len(a.status.Items), first+pageSize)
	copy, _ := a.status.ShallowCopy().(*types.StatusContainer)
	copy.Items = a.status.Items[first:last]
	return copy
}
