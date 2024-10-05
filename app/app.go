package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Since we need App.ctx to display a dialog and we can only get it when Startup method
// is executed, we keep track of the first fatal error that has happened before Startup
var startupError error

// Find: NewViews
type App struct {
	ctx context.Context

	session    config.Session
	apiKeys    map[string]string
	renderCtxs map[base.Address][]*output.RenderCtx
	meta       coreTypes.MetaData
	globals    sdk.Globals

	// Summaries
	abis              types.AbiContainer
	index             types.IndexContainer
	manifest          types.ManifestContainer
	monitors          types.MonitorContainer
	names             types.NameContainer
	status            types.StatusContainer
	project           types.ProjectContainer
	ScraperController *daemons.DaemonScraper
	FreshenController *daemons.DaemonFreshen
	IpfsController    *daemons.DaemonIpfs
}

// Find: NewViews
func NewApp() *App {
	a := App{
		apiKeys:    make(map[string]string),
		renderCtxs: make(map[base.Address][]*output.RenderCtx),
	}
	a.monitors.MonitorMap = make(map[base.Address]coreTypes.Monitor)
	a.names.NamesMap = make(map[base.Address]coreTypes.Name)
	a.project = types.NewProjectContainer("Untitled.tbx", &types.HistoryMap{}, &sync.Map{}, &sync.Map{})

	// it's okay if it's not found
	a.session.MustLoadSession()
	a.globals = sdk.Globals{
		Chain: a.session.Chain,
	}

	if err := godotenv.Load(); err != nil {
		// a.Fatal("Error loading .env file")
		logger.Info("Could not load .env file") // we don't need it for this app
		// } else if a.apiKeys["openAi"] = os.Getenv("OPENAI_API_KEY"); a.apiKeys["openAi"] == "" {
		// 	log.Fatal("No OPENAI_API_KEY key found")
	}

	// Initialize your data here

	return &a
}

func (a *App) String() string {
	bytes, _ := json.MarshalIndent(a, "", "  ")
	return string(bytes)
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

// Find: NewViews
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	a.FreshenController = daemons.NewFreshen(a, "freshen", 3000, a.GetSessionDeamon("daemon-freshen"))
	a.ScraperController = daemons.NewScraper(a, "scraper", 7000, a.GetSessionDeamon("daemon-scraper"))
	a.IpfsController = daemons.NewIpfs(a, "ipfs", 10000, a.GetSessionDeamon("daemon-ipfs"))
	go a.startDaemons()

	if startupError != nil {
		a.Fatal(startupError.Error())
	}

	logger.Info("Starting freshen process...")
	a.Refresh(a.GetSession().LastRoute)

	if err := a.loadConfig(); err != nil {
		messages.SendError(a.ctx, err)
	}

	addr := strings.ReplaceAll(a.GetSessionSubVal("/history"), "/", "")
	if len(addr) > 0 {
		logger.Info("Loading history for address: ", addr)
		go a.HistoryPage(addr, -1, 15)
	}
}

func (a *App) DomReady(ctx context.Context) {
	runtime.WindowSetPosition(a.ctx, a.session.Window.X, a.session.Window.Y)
	runtime.WindowSetSize(a.ctx, a.session.Window.Width, a.session.Window.Height)
	runtime.WindowShow(a.ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	a.session.Window.X, a.session.Window.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Window.Width, a.session.Window.Height = runtime.WindowGetSize(a.ctx)
	a.session.Window.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	a.session.Save()
}

func (a *App) GetSession() *config.Session {
	if a.session.LastSub == nil {
		a.session.LastSub = make(map[string]string)
	}
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

func (a *App) GetEnv(key string) string {
	return os.Getenv(key)
}

func (a *App) SetEnv(key, value string) {
	os.Setenv(key, value)
}

func (a *App) GetMeta() coreTypes.MetaData {
	return a.meta
}

type ModifyData struct {
	Operation string       `json:"operation"`
	Address   base.Address `json:"address"`
	Value     string       `json:"value"`
}

func (a *App) ModifyNoop(modData *ModifyData) error {
	route := a.GetSessionVal("route")
	messages.Send(a.ctx, messages.Info, messages.NewInfoMessage(fmt.Sprintf("%s modify %s: %s", route, modData.Operation, modData.Address.Hex())))
	return nil
}
