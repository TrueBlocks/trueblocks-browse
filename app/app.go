package app

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ---------------------------------------------------------------
type App struct {
	ctx context.Context
	cfg coreConfig.ConfigFile

	session config.Session

	meta    coreTypes.MetaData
	globals sdk.Globals

	renderCtxs map[base.Address][]*output.RenderCtx

	// Containers
	abis     types.AbiContainer
	index    types.IndexContainer
	manifest types.ManifestContainer
	monitors types.MonitorContainer
	names    types.NameContainer
	status   types.StatusContainer
	project  types.ProjectContainer

	// Controllers
	ScraperController *daemons.DaemonScraper
	FreshenController *daemons.DaemonFreshen
	IpfsController    *daemons.DaemonIpfs
}

// ---------------------------------------------------------------
func NewApp() *App {
	a := App{
		renderCtxs: make(map[base.Address][]*output.RenderCtx),
	}
	a.names.NamesMap = make(map[base.Address]coreTypes.Name)
	a.project = types.NewProjectContainer("Untitled.tbx", &types.HistoryMap{}, &sync.Map{}, &sync.Map{})
	a.session.LastSub = make(map[string]string)

	return &a
}

// ---------------------------------------------------------------
func (a *App) String() string {
	bytes, _ := json.MarshalIndent(a, "", "  ")
	return string(bytes)
}

// ---------------------------------------------------------------
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	path := "/Users/jrush/Library/Application Support/TrueBlocks/trueBlocks.toml"
	if err := coreConfig.ReadToml(path, &a.cfg); err != nil {
		messages.EmitError(a.ctx, err)
	}

	a.loadSession()

	go a.loadHistory(a.GetLastAddress(), nil, nil)

	a.FreshenController = daemons.NewFreshen(a, "freshen", 3000, a.GetSessionDeamon("daemon-freshen"))
	a.ScraperController = daemons.NewScraper(a, "scraper", 7000, a.GetSessionDeamon("daemon-scraper"))
	a.IpfsController = daemons.NewIpfs(a, "ipfs", 10000, a.GetSessionDeamon("daemon-ipfs"))
	go a.startDaemons()

	logger.Info("Starting freshen process...")
	a.Refresh(a.session.LastRoute)
}

// ---------------------------------------------------------------
func (a *App) DomReady(ctx context.Context) {
	win := a.GetWindow()
	runtime.WindowSetPosition(a.ctx, win.X, win.Y)
	runtime.WindowSetSize(a.ctx, win.Width, win.Height)
	runtime.WindowShow(a.ctx)
}

// ---------------------------------------------------------------
func (a *App) Shutdown(ctx context.Context) {
	a.saveSession()
}

// ---------------------------------------------------------------
func (a *App) GetSession() *config.Session {
	return &a.session
}

// ---------------------------------------------------------------
func (a *App) GetContext() context.Context {
	return a.ctx
}

// ---------------------------------------------------------------
func (a *App) GetEnv(key string) string {
	return os.Getenv(key)
}

// ---------------------------------------------------------------
func (a *App) SetEnv(key, value string) {
	os.Setenv(key, value)
}

// ---------------------------------------------------------------
func (a *App) GetMeta() coreTypes.MetaData {
	return a.meta
}

// ---------------------------------------------------------------
func (a *App) GetWindow() *config.Window {
	return &a.session.Window
}

// ---------------------------------------------------------------
func (a *App) saveSession() {
	a.session.Window.X, a.session.Window.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Window.Width, a.session.Window.Height = runtime.WindowGetSize(a.ctx)
	a.session.Window.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	a.session.Save()
}

// ----------------------------------------------------------------
func (a *App) loadSession() {
	a.session.Load()
	a.globals = sdk.Globals{
		Chain: a.session.Chain,
	}
}
