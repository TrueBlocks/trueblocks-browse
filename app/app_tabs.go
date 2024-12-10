// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// JIMMY_JAM

// JIMMY_JAM

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

	route := a.GetLastRoute()
	if tab, ok := tabMap[route]; ok {
		return tab
	}
	return []string{"no-name"}
}
