package app

import (
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
	// freshenMutex.Lock()
	// a.RefreshPaused = true
	a.globals.Chain = chain
	a.SetSessionVal("chain", chain)
	// a.CancleContexts()
	// a.ensMap = make(map[string]base.Address)
	// a.renderCtxs = make(map[base.Address][]*output.RenderCtx)
	// a.historyMap = make(map[base.Address]types.HistoryContainer)
	// a.balanceMap = sync.Map{}
	// a.meta = coreTypes.MetaData{}
	// a.abis = types.AbiContainer{}
	// a.index = types.IndexContainer{}
	// a.manifest = types.ManifestContainer{}
	// a.monitors = types.MonitorContainer{}
	// a.names = types.NameContainer{}
	// a.status = types.StatusContainer{}
	// a.portfolio = types.PortfolioContainer{}
	// a.RefreshPaused = false
	// freshenMutex.Unlock()
	// a.Reload(address)
}
