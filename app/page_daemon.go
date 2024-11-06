// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) DaemonPage(first, pageSize int) *types.DaemonContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	copy, _ := a.daemons.ShallowCopy().(*types.DaemonContainer)
	return copy
}
