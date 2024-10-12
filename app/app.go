package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ---------------------------------------------------------------
type App struct {
	ctx context.Context

	session    config.Session
	renderCtxs map[base.Address][]*output.RenderCtx
	meta       coreTypes.MetaData
	globals    sdk.Globals

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

	// it's okay if it's not found
	a.session.MustLoadSession()
	a.globals = sdk.Globals{
		Chain: a.session.Chain,
	}

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

	if err := a.loadConfig(); err != nil {
		messages.SendError(a.ctx, err)
	}

	go a.loadHistory(a.GetLastAddress(), nil, nil)

	a.FreshenController = daemons.NewFreshen(a, "freshen", 3000, a.GetSessionDeamon("daemon-freshen"))
	a.ScraperController = daemons.NewScraper(a, "scraper", 7000, a.GetSessionDeamon("daemon-scraper"))
	a.IpfsController = daemons.NewIpfs(a, "ipfs", 10000, a.GetSessionDeamon("daemon-ipfs"))
	go a.startDaemons()

	logger.Info("Starting freshen process...")
	a.Refresh(a.GetSession().LastRoute)
}

// ---------------------------------------------------------------
func (a *App) DomReady(ctx context.Context) {
	runtime.WindowSetPosition(a.ctx, a.session.Window.X, a.session.Window.Y)
	runtime.WindowSetSize(a.ctx, a.session.Window.Width, a.session.Window.Height)
	runtime.WindowShow(a.ctx)
}

// ---------------------------------------------------------------
func (a *App) Shutdown(ctx context.Context) {
	a.session.Window.X, a.session.Window.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Window.Width, a.session.Window.Height = runtime.WindowGetSize(a.ctx)
	a.session.Window.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	a.session.Save()
}

// ---------------------------------------------------------------
func (a *App) GetSession() *config.Session {
	if a.session.LastSub == nil {
		a.session.LastSub = make(map[string]string)
	}
	return &a.session
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
func (a *App) GetContext() context.Context {
	return a.ctx
}
