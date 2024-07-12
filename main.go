package main

import (
	"embed"
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	a := app.NewApp()
	opts := options.App{
		Title:            "ApplicationTitle",
		Width:            a.GetSession().Width,
		Height:           a.GetSession().Height,
		OnStartup:        a.Startup,
		OnDomReady:       a.DomReady,
		OnShutdown:       a.Shutdown,
		BackgroundColour: nil,
		Bind: []interface{}{
			a,
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
