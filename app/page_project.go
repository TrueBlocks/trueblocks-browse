// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) ProjectPage(first, pageSize int) *types.ProjectContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.project.Items)-1))
	last := base.Min(len(a.project.Items), first+pageSize)
	copy, _ := a.project.ShallowCopy().(*types.ProjectContainer)
	copy.Items = a.project.Items[first:last]
	return copy
}
