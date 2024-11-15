package app

import (
	"context"
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
	filterMap    *types.FilterMap
	namesMap     map[base.Address]coreTypes.Name
	historyCache *types.HistoryMap
	renderCtxs   map[base.Address][]*output.RenderCtx

	// Controllers
	scraperController *daemons.DaemonScraper
	freshenController *daemons.DaemonFreshen
	ipfsController    *daemons.DaemonIpfs

	timer Timer
}

func NewApp() *App {
	a := &App{
		ensCache:     &sync.Map{},
		balanceCache: &sync.Map{},
		filterMap:    &types.FilterMap{},
		namesMap:     make(map[base.Address]coreTypes.Name),
		historyCache: &types.HistoryMap{},
		renderCtxs:   make(map[base.Address][]*output.RenderCtx),
	}
	a.session.LastSub = make(map[string]string)
	a.timer = NewTimer()

	return a
}

func (a *App) Startup(ctx context.Context) {
	defer a.trackPerformance("Startup", false)()
	a.ctx = ctx
}

// DomReady is called by Wails when the app is ready to go. Adjust the window size and show it.
func (a *App) DomReady(ctx context.Context) {
	defer a.trackPerformance("DomReady", false)()

	// This call does a number of things. If any errors occur, they are deferred until
	// the window is open. This is because we can't show errors until the window is open.
	// The process is:
	// 1. Loads the session file (session.json)
	// 2. Loads the configuration file (trueBlocks.toml)
	// 3. Pings the rpcProvider (read from config file)
	// 4. If ping works, loads the names database
	// 5. If loading the names database works, starts the daemons
	// 6. In any case, makes sure the window is positioned and sized (even if all others fail)
	_ = a.initialize()

	// A properly sized window is always ready to show even if there were errors...
	runtime.WindowShow(a.ctx)

	// Now that the window is opened...
	if a.cntWizErrs() > 0 {

		// ...show any error (if there are any)...
		a.emitWizErrs()
		a.setWizardState(types.WizWelcome)
		logger.Info("There were errors during initialization...")

	} else {
		// we are initialized sucessfully, so load the latest project file
		// and freshen it.
		fn := a.getFullPath()
		if file.FileExists(fn) {
			a.readFile(fn)
		} else {
			a.newFile()
		}

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
		Chain: a.getChain(),
		// Output:  a.Output,
		// Append:  a.Append,
	}
}

/*
Here’s a summary of the decisions we made for each feature, organized by category:

Included Features
These functionalities are actively included in your function.

Search: Allows locating specific items within a container based on defined keywords or criteria.
Sort: Orders items in ascending or descending order based on specified attributes.
Filter: Displays only items that meet certain conditions for focused analysis.
Aggregate: Summarizes data by performing operations like sum, average, or count on grouped items.
Group: Organizes items into categories or groups, such as by month or year, for comparative analysis.
Transform: Modifies item attributes to new forms, supporting data normalization or enrichment.
Excluded Features
These functionalities are permanently excluded.

Export: Allows exporting processed data into formats like CSV or JSON for external use.
Visualize: Generates visual data representations, such as charts or graphs.
Audit: Tracks changes or access patterns within the data for security and compliance.
Potential Future Upgrades
These are excluded for now but could be added later if needed.

Validate: Checks items against rules or schemas to ensure data integrity.
Merge: Combines data from multiple containers or sources for comprehensive analysis.
Annotate: Adds metadata or notes to items, providing context for future reference.
This structure provides flexibility for handling data while keeping future expansions in mind. Let me know if there’s anything else you’d like to adjust or explore further!
*/
