package app

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func (a *App) GetMenus() *menu.Menu {
	appMenu := menu.NewMenu()

	systemMenu := appMenu.AddSubmenu("System")
	systemMenu.AddText("About", nil, a.SystemAbout)
	systemMenu.AddText("Quit", keys.CmdOrCtrl("q"), a.SystemQuit)

	fileMenu := appMenu.AddSubmenu("File")
	fileMenu.AddText("New File", keys.CmdOrCtrl("n"), a.FileNew)
	fileMenu.AddText("Open File", keys.CmdOrCtrl("o"), a.FileOpen)
	fileMenu.AddText("Save File", keys.CmdOrCtrl("s"), a.FileSave)
	fileMenu.AddText("Save As File", nil, a.FileSaveAs)

	viewMenu := appMenu.AddSubmenu("View")
	viewMenu.AddText("Project", keys.CmdOrCtrl("1"), a.ProjectView)
	viewMenu.AddText("History", keys.CmdOrCtrl("2"), a.HistoryView)
	viewMenu.AddText("Monitors", keys.CmdOrCtrl("3"), a.MonitorsView)
	viewMenu.AddText("Names", keys.CmdOrCtrl("4"), a.NamesView)
	viewMenu.AddText("Abis", keys.CmdOrCtrl("5"), a.AbisView)
	viewMenu.AddText("Indexes", keys.CmdOrCtrl("6"), a.IndexesView)
	viewMenu.AddText("Manifest", keys.CmdOrCtrl("7"), a.ManifestsView)
	viewMenu.AddText("Status", keys.CmdOrCtrl("8"), a.StatusView)
	viewMenu.AddText("Settings", keys.CmdOrCtrl("9"), a.SettingsView)
	viewMenu.AddText("Daemons", keys.CmdOrCtrl("0"), a.DaemonsView)
	viewMenu.AddText("Wizard", keys.CmdOrCtrl("w"), a.WizardView)

	helpMenu := appMenu.AddSubmenu("Help")
	helpMenu.AddText("Show Header", keys.CmdOrCtrl("e"), a.HeaderToggle)
	helpMenu.AddText("Show Menu", keys.CmdOrCtrl("m"), a.MenuToggle)
	helpMenu.AddText("Show Help", keys.CmdOrCtrl("h"), a.HelpToggle)
	helpMenu.AddText("Show Footer", keys.CmdOrCtrl("f"), a.FooterToggle)
	helpMenu.AddText("Show Accordion", keys.CmdOrCtrl("x"), a.AccordionToggle)
	helpMenu.AddText("Prev Tab", keys.Combo("n", keys.CmdOrCtrlKey, keys.ShiftKey), a.SwitchTabPrev)
	helpMenu.AddText("Next Tab", keys.Combo("p", keys.CmdOrCtrlKey, keys.ShiftKey), a.SwitchTabNext)

	return appMenu
}
