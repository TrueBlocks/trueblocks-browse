package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetChain() string {
	return a.globals.Chain
}

func (a *App) GetChainList() []string {
	return []string{
		"mainnet",
		"sepolia",
		"gnosis",
	}
}

func (a *App) SetChain(chain string, address base.Address) {
	a.CancelContexts() // cancel what's happening on the old chain
	a.globals.Chain = chain
	a.SetSessionVal("chain", chain)
	a.Reload(address)
	a.abis = types.AbiContainer{}
	a.index = types.IndexContainer{}
	a.manifest = types.ManifestContainer{}
	a.monitors = types.MonitorContainer{}
	a.Refresh(false)
}
