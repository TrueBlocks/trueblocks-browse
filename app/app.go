package app

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
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
	Documents  []types.Document
	CurrentDoc *types.Document

	session    config.Session
	apiKeys    map[string]string
	ensMap     map[string]base.Address
	renderCtxs map[base.Address][]*output.RenderCtx
	historyMap map[base.Address]types.SummaryTransaction
	balanceMap map[base.Address]string

	// Summaries
	abis     types.AbiContainer
	index    types.IndexContainer
	manifest types.ManifestContainer
	monitors types.MonitorContainer
	names    types.NameContainer
	status   types.SummaryStatus

	// Add your application's data here
	ScraperController *daemons.DaemonScraper
	FreshenController *daemons.DaemonFreshen
	IpfsController    *daemons.DaemonIpfs
}

// Find: NewViews
func NewApp() *App {
	a := App{
		apiKeys:    make(map[string]string),
		renderCtxs: make(map[base.Address][]*output.RenderCtx),
		ensMap:     make(map[string]base.Address),
		// Initialize maps here
		historyMap: make(map[base.Address]types.SummaryTransaction),
		balanceMap: make(map[base.Address]string),
		Documents:  make([]types.Document, 10),
	}
	a.monitors.MonitorMap = make(map[base.Address]coreTypes.Monitor)
	a.names.NamesMap = make(map[base.Address]coreTypes.Name)
	a.CurrentDoc = &a.Documents[0]
	a.CurrentDoc.Filename = "Untitled"

	// it's okay if it's not found
	_ = a.session.Load()

	if err := godotenv.Load(); err != nil {
		// a.Fatal("Error loading .env file")
		logger.Info("Could not load .env file") // we don't need it for this app
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

func (a *App) GetContext() context.Context {
	return a.ctx
}

var freshenLock atomic.Uint32

// Freshen gets called by the daemons to instruct first the backend, then the frontend to update.
// Protect against updating too fast... Note that this routine is called as a goroutine.
func (a *App) Freshen(which ...string) {
	// Skip this update we're actively upgrading
	if !freshenLock.CompareAndSwap(0, 1) {
		// logger.Info(colors.Red, "Skipping update", colors.Off)
		return
	}
	logger.Info(colors.Green, "Freshening...", colors.Off)
	defer freshenLock.CompareAndSwap(1, 0)

	notify :=
		func() {
			// Let the front end know it needs to update
			messages.Send(a.ctx, messages.Daemon, messages.NewDaemonMsg(
				a.FreshenController.Color,
				"Freshening...",
				a.FreshenController.Color,
			))
		}

	// First, we want to update the current route if we're told to
	route := ""
	if len(which) > 0 {
		route = which[0]
	}
	switch route {
	case "/abis":
		a.loadAbis(nil)
		notify()
	case "/manifest":
		a.loadManifest(nil)
		notify()
	case "/monitors":
		a.loadMonitors(nil)
		notify()
	case "/names":
		a.loadNames(nil)
		notify()
	case "/index":
		a.loadIndex(nil)
		notify()
	}

	// Now update everything in the fullness of time
	wg := sync.WaitGroup{}
	wg.Add(5)
	go a.loadAbis(&wg)
	go a.loadManifest(&wg)
	go a.loadMonitors(&wg)
	go a.loadNames(&wg)
	go a.loadIndex(&wg)
	wg.Wait()
	notify()
}

// Find: NewViews
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	a.FreshenController = daemons.NewFreshen(a, "freshen", 2000, a.GetLastDaemon("daemon-freshen"))
	a.ScraperController = daemons.NewScraper(a, "scraper", 7000, a.GetLastDaemon("daemon-scraper"))
	a.IpfsController = daemons.NewIpfs(a, "ipfs", 10000, a.GetLastDaemon("daemon-ipfs"))
	go a.startDaemons()

	if startupError != nil {
		a.Fatal(startupError.Error())
	}

	logger.Info("Starting freshen process...")
	a.Freshen(a.GetSession().LastRoute)

	// now = time.Now()
	if err := a.loadStatus(); err != nil {
		logger.Panic(err)
	}
	// fmt.Println(colors.BrightYellow, "Startup time status:", time.Since(now), colors.Off)

	// now = time.Now()
	if err := a.loadConfig(); err != nil {
		logger.Panic(err)
	}
	// fmt.Println(colors.BrightYellow, "Startup time config:", time.Since(now), colors.Off)
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
