package app

import (
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) GetChain() string {
	return a.globals.Chain
}

func (a *App) GetChains() []string {
	ret := []string{}
	for _, chain := range a.cfg.Chains {
		ret = append(ret, chain.Chain)
	}
	sort.Strings(ret)
	return ret
}

func (a *App) GetChainInfo(chain string) coreTypes.Chain {
	for _, ch := range a.status.Chains {
		if ch.Chain == chain {
			return ch
		}
	}
	return coreTypes.Chain{}
}

func (a *App) SetChain(chain string, address base.Address) {
	a.CancelAllContexts() // cancel what's happening on the old chain
	a.globals.Chain = chain
	a.SetSessionVal("chain", chain)
	a.Reload(address)
	a.abis = types.AbiContainer{}
	a.index = types.IndexContainer{}
	a.manifest = types.ManifestContainer{}
	a.monitors = types.MonitorContainer{}
	a.Refresh()
}
