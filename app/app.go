package app

import (
	"context"
	"encoding/json"
	"path/filepath"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	sdk.Globals `json:",inline"`

	ctx      context.Context
	meta     coreTypes.MetaData
	Filename string

	renderCtxs map[base.Address][]*output.RenderCtx

	// Containers
	dashboard types.DashboardContainer
	monitors  types.MonitorContainer
	names     types.NameContainer
	abis      types.AbiContainer
	indexes   types.IndexContainer
	manifests types.ManifestContainer
	status    types.StatusContainer
	settings  types.SettingsGroup
	session   types.SessionContainer
	config    types.ConfigContainer

	// Memory caches
	ProjectCache *types.ProjectMap `json:"projectCache"`
	HistoryCache *types.HistoryMap `json:"historyCache"`
	EnsCache     *sync.Map         `json:"ensCache"`
	BalanceCache *sync.Map         `json:"balanceCache"`

	// Controllers
	ScraperController *daemons.DaemonScraper
	FreshenController *daemons.DaemonFreshen
	IpfsController    *daemons.DaemonIpfs
}

func NewApp() *App {
	a := App{
		renderCtxs: make(map[base.Address][]*output.RenderCtx),
	}
	a.ProjectCache = &types.ProjectMap{}
	a.HistoryCache = &types.HistoryMap{}
	a.EnsCache = &sync.Map{}
	a.BalanceCache = &sync.Map{}
	a.names.NamesCache = make(map[base.Address]coreTypes.Name)
	a.session.LastSub = make(map[string]string)

	return &a
}

func (a *App) String() string {
	bytes, _ := json.MarshalIndent(a, "", "  ")
	return string(bytes)
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.loadSession()
	a.Filename = filepath.Join(a.session.LastFolder, a.session.LastFile)
	if !file.FileExists(a.Filename) {
		a.Filename = ""
	}
	a.Chain = a.session.LastChain
	a.dashboard = types.NewDashboardContainer(a.Chain, []types.ProjectContainer{
		types.NewProjectContainer(a.Chain, []base.Address{base.HexToAddress(a.session.LastSub["/history"])}),
	})

	go a.loadHistory(a.GetAddress(), nil, nil)

	a.FreshenController = daemons.NewFreshen(a, "freshen", 3000, a.IsShowing("freshen"))
	a.ScraperController = daemons.NewScraper(a, "scraper", 7000, a.IsShowing("scraper"))
	a.IpfsController = daemons.NewIpfs(a, "ipfs", 10000, a.IsShowing("ipfs"))
	go a.startDaemons()

	logger.Info("Starting freshen process...")
	_ = a.Refresh()
}

func (a *App) DomReady(ctx context.Context) {
	win := a.GetWindow()
	runtime.WindowSetPosition(a.ctx, win.X, win.Y)
	runtime.WindowSetSize(a.ctx, win.Width, win.Height)
	runtime.WindowShow(a.ctx)

	if path, err := utils.GetConfigFn("", "trueBlocks.toml"); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
		})
	} else {
		if err := coreConfig.ReadToml(path, &a.config.Config); err != nil {
			messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
				String1: err.Error(),
			})
		}
	}
}

func (a *App) Shutdown(ctx context.Context) {
	a.saveSession()
}

func (a *App) saveSession() {
	a.session.Window.X, a.session.Window.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Window.Width, a.session.Window.Height = runtime.WindowGetSize(a.ctx)
	a.session.Window.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	_ = a.session.Save()
}

func (a *App) loadSession() {
	_ = a.session.Load()
	a.session.CleanWindowSize(a.ctx)
	a.Chain = a.session.LastChain
}

func (a *App) Logger(msg string) {
	logger.Info(msg)
}
