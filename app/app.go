package app

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
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
	ctx        context.Context
	session    config.Session
	apiKeys    map[string]string
	ensMap     map[string]base.Address
	renderCtxs map[base.Address][]*output.RenderCtx
	// Summaries
	names    types.SummaryName
	monitors types.SummaryMonitor
	index    types.SummaryIndex
	manifest types.SummaryManifest
	abis     types.SummaryAbis
	status   types.StatusSummary
	// Add your application's data here
	ScraperController *daemons.DaemonScraper
	FileController    *daemons.DaemonFile
	FreshenController *daemons.DaemonFreshen
	IpfsController    *daemons.DaemonIpfs
	Documents         []types.Document
	CurrentDoc        *types.Document
}

// Find: NewViews
func NewApp() *App {
	a := App{
		apiKeys:    make(map[string]string),
		renderCtxs: make(map[base.Address][]*output.RenderCtx),
		ensMap:     make(map[string]base.Address),
		// Initialize maps here
		ScraperController: daemons.NewScraper("scraper", 1000), // TODO: Should be seven seconds
		FileController:    daemons.NewFileDaemon("filedaemon", 8080, 1000),
		FreshenController: daemons.NewFreshen("freshen", 1000),
		IpfsController:    daemons.NewIpfs("ipfs", 1000),
		Documents:         make([]types.Document, 10),
	}
	a.monitors.MonitorMap = make(map[base.Address]coreTypes.Monitor)
	a.names.NamesMap = make(map[base.Address]coreTypes.Name)
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

// Find: NewViews
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.ScraperController.MsgCtx = ctx
	a.FileController.MsgCtx = ctx
	a.FreshenController.MsgCtx = ctx
	a.IpfsController.MsgCtx = ctx
	if startupError != nil {
		a.Fatal(startupError.Error())
	}
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
