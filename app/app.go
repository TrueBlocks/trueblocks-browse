package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx   context.Context
	meta  coreTypes.MetaData
	dirty bool

	// Containers
	project   types.ProjectContainer
	monitors  types.MonitorContainer
	names     types.NameContainer
	abis      types.AbiContainer
	indexes   types.IndexContainer
	manifests types.ManifestContainer
	status    types.StatusContainer
	settings  types.SettingsContainer
	session   types.SessionContainer
	config    types.ConfigContainer
	wizard    types.WizardContainer
	daemons   types.DaemonContainer

	// Memory caches
	ensCache     *sync.Map
	balanceCache *sync.Map
	namesMap     map[base.Address]coreTypes.Name
	historyCache *types.HistoryMap
	renderCtxs   map[base.Address][]*output.RenderCtx

	// Controllers
	scraperController *daemons.DaemonScraper
	freshenController *daemons.DaemonFreshen
	ipfsController    *daemons.DaemonIpfs

	// During initialization, we do things that may cause errors, but
	// we have not yet opened the window, so we defer them until we can
	// decide what to do.
	deferredErrors []error
}

func NewApp() *App {
	a := &App{
		ensCache:     &sync.Map{},
		balanceCache: &sync.Map{},
		namesMap:     make(map[base.Address]coreTypes.Name),
		historyCache: &types.HistoryMap{},
		renderCtxs:   make(map[base.Address][]*output.RenderCtx),
	}
	a.session.LastSub = make(map[string]string)

	return a
}

var ErrLoadingNames = errors.New("error loading names")
var ErrWindowSize = errors.New("error fixing window size")

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	// We do various setups prior to showing the window. Improves interactivity.
	// But...something may fail, so we need to keep track of errors.
	var err error
	if err = a.session.Load(); err != nil {
		a.deferredErrors = append(a.deferredErrors, err)
	}

	// Load the trueBlocks.toml file
	if err = a.config.Load(); err != nil {
		a.deferredErrors = append(a.deferredErrors, err)
	}
	if a.session.LastChain, err = a.config.IsValidChain(a.session.LastChain); err != nil {
		a.deferredErrors = append(a.deferredErrors, err)
	}

	// We always need names, so let's load it before showing the window
	if a.namesMap, err = names.LoadNamesMap(namesChain, coreTypes.All, nil); err == nil {
		wErr := fmt.Errorf("%w: %v", ErrLoadingNames, err)
		a.deferredErrors = append(a.deferredErrors, wErr)
	}

	freshenRate := time.Duration(3000)
	if os.Getenv("TB_FRESHEN_RATE") != "" {
		rate := base.MustParseInt64(os.Getenv("TB_FRESHEN_RATE"))
		if rate > 0 {
			freshenRate = time.Duration(rate)
		}
	}
	a.freshenController = daemons.NewFreshen(a, "freshen", freshenRate, a.IsShowing("freshen"))
	a.scraperController = daemons.NewScraper(a, "scraper", 7000, a.IsShowing("scraper"))
	a.ipfsController = daemons.NewIpfs(a, "ipfs", 10000, a.IsShowing("ipfs"))
}

// DomReady is called by Wails when the app is ready to go. Adjust the window size and show it.
func (a *App) DomReady(ctx context.Context) {
	var err error

	// We're ready to open the window, but first we need to make sure it will show...
	if a.session.Window, err = a.session.CleanWindowSize(a.ctx); err != nil {
		wErr := fmt.Errorf("%w: %v", ErrWindowSize, err)
		a.deferredErrors = append(a.deferredErrors, wErr)
	}
	// DO NOT COLLAPSE - A VALID WINDOW IS RETURNED EVEN ON ERROR
	runtime.WindowSetPosition(a.ctx, a.session.Window.X, a.session.Window.Y)
	runtime.WindowSetSize(a.ctx, a.session.Window.Width, a.session.Window.Height)
	runtime.WindowShow(a.ctx)
	if err != nil {
		a.deferredErrors = append(a.deferredErrors, err)
	}

	// We now have a window, so we can finally show any accumulated errors
	for _, err := range a.deferredErrors {
		a.emitErrorMsg(err, nil)
	}

	fn := a.getFullPath()
	if file.FileExists(fn) {
		a.readFile(fn)
	} else {
		a.newFile()
	}

	go a.Freshen()

	go a.freshenController.Run()
	go a.scraperController.Run()
	go a.ipfsController.Run()

	logger.Info("Fininished loading...")
}

// Shutdown is called by Wails when the app is closed
func (a *App) Shutdown(ctx context.Context) {
	a.saveSession()
}

func (a *App) getGlobals() sdk.Globals {
	return sdk.Globals{
		// Ether:   a.Ether,
		// Cache:   a.Cache,
		// Decache: a.Decache,
		// Verbose: a.Verbose,
		Chain: a.session.LastChain,
		// Output:  a.Output,
		// Append:  a.Append,
	}
}
