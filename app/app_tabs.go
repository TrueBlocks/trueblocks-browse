// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
// EXISTING_CODE

func (a *App) GetTabs() []string {
	tabMap := map[string][]string{
		"project":   {"project"},
		"history":   {"balances", "incoming", "outgoing", "internals", "charts", "logs", "statements", "neighbors", "traces", "receipts"},
		"monitors":  {"monitors"},
		"sharing":   {"names", "abis", "pins", "uploads"},
		"unchained": {"indexes", "manifests", "pins", "uploads"},
		"settings":  {"status", "config", "session"},
		"daemons":   {"daemons"},
		"wizard":    {"wizard"},
	}

	route := a.getLastRoute()
	if tab, ok := tabMap[route]; ok {
		return tab
	}
	return []string{"no-name"}
}

func (a *App) getTabs() []string {
	return a.GetTabs()
}
