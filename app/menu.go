// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

var keyMap = map[string]*keys.Accelerator{
	"dashboard": keys.CmdOrCtrl("1"),
	"history":   keys.CmdOrCtrl("2"),
	"monitors":  keys.CmdOrCtrl("3"),
	"names":     keys.CmdOrCtrl("4"),
	"abis":      keys.CmdOrCtrl("5"),
	"indexes":   keys.CmdOrCtrl("6"),
	"manifests": keys.CmdOrCtrl("7"),
	"status":    keys.CmdOrCtrl("8"),
	"settings":  keys.CmdOrCtrl("9"),
	"daemons":   keys.CmdOrCtrl("0"),
	"session":   keys.CmdOrCtrl("u"),
	"config":    keys.CmdOrCtrl("v"),
	"wizard":    keys.CmdOrCtrl("w"),
}

// EXISTING_CODE

func (a *App) GetMenus() *menu.Menu {
	// EXISTING_CODE
	appMenu := menu.NewMenu()
	systemMenu := appMenu.AddSubmenu("System")
	fileMenu := appMenu.AddSubmenu("File")
	viewMenu := appMenu.AddSubmenu("View")
	helpMenu := appMenu.AddSubmenu("Help")

	systemMenu.AddText("About", nil, a.SystemAbout)
	systemMenu.AddText("Quit", keys.CmdOrCtrl("q"), a.SystemQuit)

	fileMenu.AddText("New File", keys.CmdOrCtrl("n"), a.FileNew)
	fileMenu.AddText("Open File", keys.CmdOrCtrl("o"), a.FileOpen)
	fileMenu.AddText("Save File", keys.CmdOrCtrl("s"), a.FileSave)
	fileMenu.AddText("Save As File", nil, a.FileSaveAs)
	// EXISTING_CODE

	viewMenu.AddText("Dashboard", keyMap["dashboard"], a.DashboardView)
	viewMenu.AddText("History", keyMap["history"], a.HistoryView)
	viewMenu.AddText("Monitors", keyMap["monitors"], a.MonitorsView)
	viewMenu.AddText("Names", keyMap["names"], a.NamesView)
	viewMenu.AddText("Abis", keyMap["abis"], a.AbisView)
	viewMenu.AddText("Indexes", keyMap["indexes"], a.IndexesView)
	viewMenu.AddText("Manifests", keyMap["manifests"], a.ManifestsView)
	viewMenu.AddText("Status", keyMap["status"], a.StatusView)
	viewMenu.AddText("Settings", keyMap["settings"], a.SettingsView)
	viewMenu.AddText("Daemons", keyMap["daemons"], a.DaemonsView)
	viewMenu.AddText("Session", keyMap["session"], a.SessionView)
	viewMenu.AddText("Config", keyMap["config"], a.ConfigView)
	viewMenu.AddText("Wizard", keyMap["wizard"], a.WizardView)

	// EXISTING_CODE
	helpMenu.AddText("Show Header", keys.CmdOrCtrl("e"), a.HeaderToggle)
	helpMenu.AddText("Show Menu", keys.CmdOrCtrl("m"), a.MenuToggle)
	helpMenu.AddText("Show Help", keys.CmdOrCtrl("h"), a.HelpToggle)
	helpMenu.AddText("Show Footer", keys.CmdOrCtrl("f"), a.FooterToggle)
	helpMenu.AddText("Show Accordion", keys.CmdOrCtrl("x"), a.AccordionToggle)
	helpMenu.AddText("Prev Tab", keys.Combo("n", keys.CmdOrCtrlKey, keys.ShiftKey), a.SwitchTabPrev)
	helpMenu.AddText("Next Tab", keys.Combo("p", keys.CmdOrCtrlKey, keys.ShiftKey), a.SwitchTabNext)
	// EXISTING_CODE

	return appMenu
}
