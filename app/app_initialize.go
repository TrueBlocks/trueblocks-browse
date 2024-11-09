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
			a.addWizErr(err)
			return false
		} else {
			// we serialize the wizard state in a session string
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
		if err := a.config.Load(); err != nil {
			a.addWizErr(err)
			return false
		} else if a.session.LastChain, err = a.config.IsValidChain(a.session.LastChain); err != nil {
			a.addWizErr(err)
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
			a.addWizErr(wErr)
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
		// We always need names. We load it here and then again if it ever changes...
		initNames := func() bool {
			var err error
			if a.namesMap, err = names.LoadNamesMap(namesChain, coreTypes.All, nil); err != nil {
				wErr := fmt.Errorf("%w: %v", ErrLoadingNames, err)
				a.addWizErr(wErr)
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
			a.addWizErr(wErr)
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
