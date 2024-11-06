// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) WizardPage(first, pageSize int) *types.WizardContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	copy, _ := a.wizard.ShallowCopy().(*types.WizardContainer)
	return copy
}
