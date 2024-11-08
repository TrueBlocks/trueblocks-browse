package app

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpc"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) initialize() bool {
	// we will be recalling this repeatedly until we have a valid configuration
	// so we always start fresh...
	var err error

	initConfig := func() { // window can open, but we need an RPC provider...which means we need a config file.
		if err = a.config.Load(); err != nil {
			a.addDeferredError(err)
		} else if a.session.LastChain, err = a.config.IsValidChain(a.session.LastChain); err != nil {
			a.addDeferredError(err)
		} else {
			logger.InfoBW("Loaded config", a.cntDeferredErrors(), "errors")
		}
	}
	initConfig()

	// The rest depends on the rpc...
	initRpc := func() bool {
		os.Setenv("TB_NO_PROVIDER_CHECK", "true")
		if err = rpc.PingRpc(a.config.Chains[a.session.LastChain].RpcProvider); err != nil {
			wErr := fmt.Errorf("%w: %v", ErrLoadingRpc, err)
			a.addDeferredError(wErr)
			os.Unsetenv("TB_NO_PROVIDER_CHECK")
			return false
		} else {
			logger.InfoBW("Connected to RPC", a.cntDeferredErrors(), "errors")
			os.Unsetenv("TB_NO_PROVIDER_CHECK")
			return true
		}
	}

	if initRpc() {
		// We always need names. We load it here and then again if it ever changes...
		initNames := func() bool {
			if a.namesMap, err = names.LoadNamesMap(namesChain, coreTypes.All, nil); err != nil {
				wErr := fmt.Errorf("%w: %v", ErrLoadingNames, err)
				a.addDeferredError(wErr)
				return false
			} else {
				logger.InfoBW("Loaded names", a.cntDeferredErrors(), "errors")
				return true
			}
		}

		if initNames() {
			// The daemons should only start if everything else is working...
			go a.startDaemons()
		}
	}

	// returns true if there are no errors...
	return a.cntDeferredErrors() == 0
}
