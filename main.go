package main

import (
	"embed"
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/app"
	"github.com/TrueBlocks/trueblocks-browse/pkg/editors"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
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
		Title:            "Browse by TrueBlocks",
		Width:            1024,
		Height:           768,
		OnStartup:        a.Startup,
		OnDomReady:       a.DomReady,
		OnShutdown:       a.Shutdown,
		BackgroundColour: nil,
		LogLevel:         wLogger.ERROR,
		Menu:             a.GetMenus(),
		Bind: []interface{}{
			a,
			&messages.MessageMsg{},
			&types.Config{},
			&editors.Name{},
			&types.Daemon{},
		},
		EnumBind: []interface{}{
			types.AllDaemonStates,
			messages.AllMessages,
			types.AllWizStates,
			types.AllSteps,
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
