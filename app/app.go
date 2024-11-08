package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
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

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// DomReady is called by Wails when the app is ready to go. Adjust the window size and show it.
func (a *App) DomReady(ctx context.Context) {
	initSession := func() {
		if err := a.session.Load(); err != nil {
			a.addDeferredError(err)
		} else {
			a.session.Window.Title = "Browse by TrueBlocks"
			logger.InfoBW("Loaded session:", a.cntDeferredErrors(), "errors")
		}
	}
	initSession()
	_ = a.initialize()

	prepareWindow := func() { // window size and placement depends on session file
		var err error
		if a.session.Window, err = a.session.CleanWindowSize(a.ctx); err != nil {
			wErr := fmt.Errorf("%w: %v", ErrWindowSize, err)
			a.addDeferredError(wErr)
		} else {
			logger.InfoBW("Window size set...")
		}
		runtime.WindowSetPosition(a.ctx, a.session.Window.X, a.session.Window.Y)
		runtime.WindowSetSize(a.ctx, a.session.Window.Width, a.session.Window.Height)
	}
	prepareWindow()

	// A properly sized window is always ready to show even if there were errors...
	runtime.WindowShow(a.ctx)

	a.addDeferredError(fmt.Errorf("error for testing purposes"))

	// Now that the window is opened, show any error (and if there are any, enter wizard mode).
	if a.cntDeferredErrors() > 0 {
		// We now have a window, so we can finally show any accumulated errors
		a.emitDeferredErrors()
		if a.getWizardState() != coreTypes.Welcome {
			a.SetWizardState(coreTypes.Error)
		}
		logger.Info("There were errors during initialization...")

	} else {
		// We're initialized, let's open the last opened file (or a new file)...
		fn := a.getFullPath()
		if file.FileExists(fn) {
			a.readFile(fn)
		} else {
			a.newFile()
		}

		// freshen the data once...daemons will take over from here...
		go a.Freshen()
		logger.Info("Fininished initializing...")
	}
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
