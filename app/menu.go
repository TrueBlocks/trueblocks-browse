// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sort"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

// EXISTING_CODE

// sugar
type cbd = menu.CallbackData

type MenuDef struct {
	number int
	menu   *menu.Menu
	title  string
	hotKey *keys.Accelerator
	action func(cb *cbd)
}

func (a *App) GetMenus() *menu.Menu {
	appMenu := menu.NewMenu()
	systemMenu := appMenu.AddSubmenu("System")
	fileMenu := appMenu.AddSubmenu("File")
	viewMenu := appMenu.AddSubmenu("View")
	helpMenu := appMenu.AddSubmenu("Help")

	id := 3000
	inc := func() int {
		id++
		return id
	}

	var menuMap = map[string]MenuDef{
		// EXISTING_CODE
		"about": {1001, systemMenu, "About", nil, a.SystemAbout},
		"quit":  {1002, systemMenu, "Quit", keys.CmdOrCtrl("q"), a.SystemQuit},

		"newFile":  {2001, fileMenu, "New File", keys.CmdOrCtrl("n"), a.FileNew},
		"openFile": {2002, fileMenu, "Open File", keys.CmdOrCtrl("o"), a.FileOpen},
		"saveFile": {2003, fileMenu, "Save File", keys.CmdOrCtrl("s"), a.FileSave},
		"saveAs":   {2004, fileMenu, "Save As", keys.Combo("s", keys.CmdOrCtrlKey, keys.ShiftKey), a.FileSaveAs},
		// EXISTING_CODE

		"project":   {inc(), viewMenu, "Project", keys.CmdOrCtrl("1"), a.ProjectView},
		"history":   {inc(), viewMenu, "History", keys.CmdOrCtrl("2"), a.HistoryView},
		"monitors":  {inc(), viewMenu, "Monitors", keys.CmdOrCtrl("3"), a.MonitorsView},
		"names":     {inc(), viewMenu, "Names", keys.CmdOrCtrl("4"), a.NamesView},
		"abis":      {inc(), viewMenu, "Abis", keys.CmdOrCtrl("5"), a.AbisView},
		"indexes":   {inc(), viewMenu, "Indexes", keys.CmdOrCtrl("6"), a.IndexesView},
		"manifests": {inc(), viewMenu, "Manifests", keys.CmdOrCtrl("7"), a.ManifestsView},
		"status":    {inc(), viewMenu, "Status", keys.CmdOrCtrl("8"), a.StatusView},
		"settings":  {inc(), viewMenu, "Settings", keys.CmdOrCtrl("9"), a.SettingsView},
		"daemons":   {inc(), viewMenu, "Daemons", keys.CmdOrCtrl("0"), a.DaemonsView},
		"session":   {inc(), viewMenu, "Session", keys.CmdOrCtrl("u"), a.SessionView},
		"config":    {inc(), viewMenu, "Config", keys.CmdOrCtrl("v"), a.ConfigView},
		"wizard":    {inc(), viewMenu, "Wizard", keys.CmdOrCtrl("w"), a.WizardView},

		// EXISTING_CODE
		"showHeader":    {4010, helpMenu, "Show Header", keys.CmdOrCtrl("e"), a.ToggleHeader},
		"showMenu":      {4020, helpMenu, "Show Menu", keys.CmdOrCtrl("m"), a.ToggleMenu},
		"showHelp":      {4030, helpMenu, "Show Help", keys.CmdOrCtrl("h"), a.ToggleHelp},
		"showFooter":    {4040, helpMenu, "Show Footer", keys.CmdOrCtrl("f"), a.ToggleFooter},
		"showAccordion": {4050, helpMenu, "Show Accordion", keys.CmdOrCtrl("x"), a.ToggleAccordion},
		"prevTab":       {4060, helpMenu, "Previous Tab", keys.Combo("n", keys.CmdOrCtrlKey, keys.ShiftKey), a.TogglePrevTab},
		"nextTab":       {4070, helpMenu, "Next Tab", keys.Combo("p", keys.CmdOrCtrlKey, keys.ShiftKey), a.ToggleNextTab},
		// EXISTING_CODE
	}

	var menus []MenuDef
	for _, menuItem := range menuMap {
		menus = append(menus, menuItem)
	}
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].number < menus[j].number
	})

	for _, item := range menus {
		item.menu.AddText(item.title, item.hotKey, item.action)
	}

	return appMenu
}

// EXISTING_CODE
// EXISTING_CODE
