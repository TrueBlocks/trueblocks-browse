package main

import (
	"embed"
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/app"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/editors"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
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
			&app.AppInfo{},
			&messages.MessageMsg{},
			&coreTypes.Wizard{},
			&coreTypes.Transaction{},
			&configTypes.Config{},
			&editors.Name{},
			&daemons.Daemon{},
		},
		EnumBind: []interface{}{
			daemons.AllStates,
			messages.AllMessages,
			coreTypes.AllStates,
			coreTypes.AllSteps,
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
