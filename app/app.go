package app

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/servers"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Since we need App.ctx to display a dialog and we can only get it when Startup method
// is executed, we keep track of the first fatal error that has happened before Startup
var startupError error

// Find: NewViews
type App struct {
	ctx         context.Context
	session     config.Session
	apiKeys     map[string]string
	namesMap    map[base.Address]types.NameEx
	names       []types.NameEx
	monitorsMap map[base.Address]types.MonitorEx
	monitors    []types.MonitorEx
	abis        []coreTypes.AbiFile
	index       types.IndexSummary
	manifest    coreTypes.Manifest
	status      coreTypes.Status
	ensMap      map[string]base.Address
	renderCtxs  map[base.Address][]*output.RenderCtx
	// Add your application's data here
	Scraper    *servers.Scraper
	FileServer *servers.FileServer
	Monitor    *servers.Monitor
	Ipfs       *servers.Ipfs
	Documents  []types.Document
	CurrentDoc *types.Document
}

// Find: NewViews
func NewApp() *App {
	a := App{
		apiKeys:     make(map[string]string),
		namesMap:    make(map[base.Address]types.NameEx),
		monitorsMap: make(map[base.Address]types.MonitorEx),
		renderCtxs:  make(map[base.Address][]*output.RenderCtx),
		ensMap:      make(map[string]base.Address),
		// Initialize maps here
		Scraper:    servers.NewScraper("scraper", 1000), // TODO: Should be seven seconds
		FileServer: servers.NewFileServer("fileserver", 8080, 1000),
		Monitor:    servers.NewMonitor("monitor", 1000),
		Ipfs:       servers.NewIpfs("ipfs", 1000),
		Documents:  make([]types.Document, 10),
	}
	a.CurrentDoc = &a.Documents[0]
	a.CurrentDoc.Filename = "Untitled"

	// it's okay if it's not found
	_ = a.session.Load()

	if err := godotenv.Load(); err != nil {
		a.Fatal("Error loading .env file")
		// } else if a.apiKeys["openAi"] = os.Getenv("OPENAI_API_KEY"); a.apiKeys["openAi"] == "" {
		// 	log.Fatal("No OPENAI_API_KEY key found")
	}

	// Initialize your data here

	return &a
}

func (a App) String() string {
	bytes, _ := json.MarshalIndent(a, "", "  ")
	return string(bytes)
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	if startupError != nil {
		a.Fatal(startupError.Error())
	}
	// Find: NewViews
	if err := a.loadNames(); err != nil {
		logger.Panic(err)
	}
	if err := a.loadMonitors(); err != nil {
		logger.Panic(err)
	}
	if err := a.loadStatus(); err != nil {
		logger.Panic(err)
	}
	if err := a.loadManifest(); err != nil {
		logger.Panic(err)
	}
	if err := a.loadAbis(); err != nil {
		logger.Panic(err)
	}
	if err := a.loadIndex(); err != nil {
		logger.Panic(err)
	}
	a.Scraper.MsgCtx = ctx
	a.FileServer.MsgCtx = ctx
	a.Monitor.MsgCtx = ctx
	a.Ipfs.MsgCtx = ctx
}

func (a *App) DomReady(ctx context.Context) {
	// Sometimes useful for debugging
	if os.Getenv("TB_CMD_LINE") == "true" {
		return
	}
	runtime.WindowSetPosition(a.ctx, a.session.X, a.session.Y)
	runtime.WindowSetSize(a.ctx, a.session.Width, a.session.Height)
	runtime.WindowShow(a.ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	// Sometimes useful for debugging
	if os.Getenv("TB_CMD_LINE") == "true" {
		return
	}
	a.session.X, a.session.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Width, a.session.Height = runtime.WindowGetSize(a.ctx)
	a.session.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	a.session.Save()
}

func (a *App) GetSession() *config.Session {
	return &a.session
}

func (a *App) Fatal(message string) {
	if message == "" {
		message = "Fatal error occured. The application cannot continue to run."
	}
	log.Println(message)

	// If a.ctx has not been set yet (i.e. we are before calling Startup), we can't display the
	// dialog. Instead, we keep the error and let Startup call this function again when a.ctx is set.
	if a.ctx == nil {
		// We will only display the first error, since it makes more sense
		if startupError == nil {
			startupError = errors.New(message)
		}
		// Return to allow the application to continue starting up, until we get the context
		return
	}
	_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Fatal Error",
		Message: message,
	})
	os.Exit(1)
}
