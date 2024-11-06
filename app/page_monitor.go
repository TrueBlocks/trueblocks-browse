// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) MonitorPage(first, pageSize int) *types.MonitorContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.monitors.Items)-1))
	last := base.Min(len(a.monitors.Items), first+pageSize)
	copy, _ := a.monitors.ShallowCopy().(*types.MonitorContainer)
	copy.Items = a.monitors.Items[first:last]
	return copy
}
