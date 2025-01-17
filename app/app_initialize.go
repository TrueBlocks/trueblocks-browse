package app

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpc"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) initialize() bool {
	initSession := func() bool {
		if err := a.session.Load(); err != nil {
			a.addWizErr(WizReasonNoSession, types.WizConfig, err)
			return false
		} else {
			// we serialize the wizard state in a session string
			// TODO: BOGUS a.wizard = types. NewWizzardContainer(a.getChain(), []types.WizError{})
			a.wizard.Chain = a.getChain()
			a.wizard.State = types.WizState(a.session.WizardStr)
			a.session.Window.Title = "Browse by TrueBlocks"
			logger.InfoBW("Loaded session:", a.cntWizErrs(), "errors")
			return true
		}
	}
	_ = initSession()

	// we will be recalling this repeatedly until we have a valid configuration
	// so we always start fresh...
	initConfig := func() bool { // window can open, but we need an RPC provider...which means we need a config file.
		a.config.Updater = types.NewConfigUpdater("unused", true)
		if err := a.loadConfig(nil, nil); err != nil {
			a.addWizErr(WizReasonNoConfig, types.WizConfig, err)
			return false
		} else if a.session.LastChain, err = a.config.IsValidChain(a.getChain()); err != nil {
			a.addWizErr(WizReasonChainNotConfigured, types.WizConfig, err)
			return false
		} else {
			logger.InfoBW("Loaded config", a.cntWizErrs(), "errors")
			return true
		}
	}
	_ = initConfig()

	// The rest depends on the rpc...
	initRpc := func() bool {
		os.Setenv("TB_NO_PROVIDER_CHECK", "true")
		if err := rpc.PingRpc(a.config.Chains[a.session.LastChain].RpcProvider); err != nil {
			wErr := fmt.Errorf("%w: %v", ErrLoadingRpc, err)
			a.addWizErr(WizReasonFailedRpcPing, types.WizRpc, wErr)
			os.Unsetenv("TB_NO_PROVIDER_CHECK")
			return false
		} else {
			logger.InfoBW("Connected to RPC", a.cntWizErrs(), "errors")
			os.Unsetenv("TB_NO_PROVIDER_CHECK")
			return true
		}
	}
	passed := initRpc()

	if passed {
		// initialize the Updaters
		initUpdaters := func() bool {
			chain := a.getChain()
			a.project.Updater = types.NewProjectUpdater(chain, []types.HistoryContainer{}, true /* reset */)
			a.monitors.Updater = types.NewMonitorUpdater(chain, true)
			a.names.Updater = types.NewNameUpdater(chain, true)
			a.abis.Updater = types.NewAbiUpdater(chain, true)
			a.indexes.Updater = types.NewIndexUpdater(chain, true)
			a.manifests.Updater = types.NewManifestUpdater(chain, true)
			a.status.Updater = types.NewStatusUpdater(chain, true)
			a.settings.Updater = types.NewSettingsUpdater(chain, true)
			a.session.Updater = types.NewSessionUpdater(chain, true)
			a.config.Updater = types.NewConfigUpdater(chain, true)
			a.wizard.Updater = types.NewWizardUpdater(chain, true)
			a.daemons.Updater = types.NewDaemonUpdater(chain, true)
			return true
		}
		_ = initUpdaters()

		// We always need names. We load it here and then again if it ever changes...
		initNames := func() bool {
			var err error
			if a.namesMap, err = names.LoadNamesMap(namesChain, coreTypes.All, nil); err != nil {
				wErr := fmt.Errorf("%w: %v", ErrLoadingNames, err)
				a.addWizErr(WizReasonFailedNamesLoad, types.WizRpc, wErr)
				return false
			} else {
				logger.InfoBW("Loaded names", a.cntWizErrs(), "errors")
				return true
			}
		}

		// The daemons should only start if everything else is working...
		passed := initNames()
		if passed {
			go a.startDaemons()
		}
	}

	prepareWindow := func() bool { // window size and placement depends on session file
		ret := false // do not collapse...we position the window below on both error and not
		var err error
		if a.session.Window, err = a.session.CleanWindowSize(a.ctx); err != nil {
			wErr := fmt.Errorf("%w: %v", ErrWindowSize, err)
			a.addWizErr(WizReasonFailedPrepareWindow, types.WizRpc, wErr)
		} else {
			logger.InfoBW("Window size set...")
			ret = true
		}
		runtime.WindowSetPosition(a.ctx, a.session.Window.X, a.session.Window.Y)
		runtime.WindowSetSize(a.ctx, a.session.Window.Width, a.session.Window.Height)
		return ret
	}
	_ = prepareWindow()

	// returns true if there are no errors...
	if a.cntWizErrs() > 0 {
		// ...goes to wizard mode and returns false otherwise
		a.setWizardState(types.WizWelcome)
		return false
	}

	return true
}
