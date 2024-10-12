package main

import (
	"embed"
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/app"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/wailsapp/wails/v2"
	wLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	a := app.NewApp()
	opts := options.App{
		Title:            a.GetSession().Window.Title,
		Width:            a.GetSession().Window.Width,
		Height:           a.GetSession().Window.Height,
		OnStartup:        a.Startup,
		OnDomReady:       a.DomReady,
		OnShutdown:       a.Shutdown,
		BackgroundColour: nil,
		LogLevel:         wLogger.ERROR,
		Menu:             a.GetMenus(),
		Bind: []interface{}{
			a,
			&messages.DocumentMsg{},
			&messages.ErrorMsg{},
			&messages.InfoMsg{},
			&messages.ProgressMsg{},
			&messages.DaemonMsg{},
			&messages.NavigateMsg{},
			&messages.HelpMsg{},
			&types.AbiContainer{},
			&types.ProjectContainer{},
			&types.IndexContainer{},
			&types.ManifestContainer{},
			&types.MonitorContainer{},
			&types.NameContainer{},
			&types.StatusContainer{},
			&wizard.Wizard{},
			&coreTypes.Transaction{},
			&daemons.Daemon{},
		},
		EnumBind: []interface{}{
			daemons.AllStates,
			messages.AllMessages,
			wizard.AllStates,
			wizard.AllSteps,
		},
		StartHidden: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
	}

	if err := wails.Run(&opts); err != nil {
		fmt.Println("Error:", err.Error())
	}
}
