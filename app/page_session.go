// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) SessionPage(first, pageSize int) *types.SessionContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	copy, _ := a.session.ShallowCopy().(*types.SessionContainer)
	return copy
}
