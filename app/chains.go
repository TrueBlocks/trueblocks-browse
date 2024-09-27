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
	a.globals.Chain = chain
	a.SetSessionVal("chain", chain)
	a.Refresh()
}
