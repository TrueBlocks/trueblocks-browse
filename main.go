package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-browse/app"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/wailsapp/wails/v2"
	wLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if os.Getenv("TB_CMD_LINE") == "true" {
		logger.Info("Running in console mode")
		a := app.NewApp()
		ctx := context.Background()
		a.Startup(ctx)
		a.DomReady(ctx)
	} else {
		a := app.NewApp()
		opts := options.App{
			Title:            a.GetSession().Title,
			Width:            a.GetSession().Width,
			Height:           a.GetSession().Height,
			OnStartup:        a.Startup,
			OnDomReady:       a.DomReady,
			OnShutdown:       a.Shutdown,
			BackgroundColour: nil,
			LogLevel:         wLogger.ERROR,
			Menu:             a.GetMenus(),
			// Find: NewViews
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
				&types.PortfolioContainer{},
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
			Debug: options.Debug{
				OpenInspectorOnStartup: a.GetEnv("TB_DEBUG") == "true",
			},
		}

		if err := wails.Run(&opts); err != nil {
			fmt.Println("Error:", err.Error())
		}
	}
}
