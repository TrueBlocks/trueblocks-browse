// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) SettingsPage(first, pageSize int) *types.SettingsContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	copy, _ := a.settings.ShallowCopy().(*types.SettingsContainer)
	return copy
}
