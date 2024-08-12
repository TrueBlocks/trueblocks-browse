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
				&messages.ProgressMsg{},
				&messages.DaemonMsg{},
				&daemons.Daemon{},
				&coreTypes.Transaction{},
				&types.AbiContainer{},
				&types.IndexContainer{},
				&types.ManifestContainer{},
				&types.MonitorContainer{},
				&types.NameContainer{},
				&types.SummaryStatus{},
			},
			EnumBind: []interface{}{
				types.NameParts,
				daemons.Types,
				daemons.States,
				messages.Messages,
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
}
