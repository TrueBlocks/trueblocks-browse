package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

var bMutex sync.Mutex

func (a *App) getBalance(address base.Address) string {
	bMutex.Lock()
	_, exists := a.balanceMap[address]
	bMutex.Unlock()

	if exists {
		bMutex.Lock()
		defer bMutex.Unlock()
		return a.balanceMap[address]
	}

	opts := sdk.StateOptions{
		Addrs: []string{address.Hex()},
		Globals: sdk.Globals{
			Ether: true,
			Cache: true,
		},
	}
	if balances, _, err := opts.State(); err != nil {
		return "0"
	} else {
		bMutex.Lock()
		defer bMutex.Unlock()
		a.balanceMap[address] = balances[0].Balance.ToEtherStr(18)
		return a.balanceMap[address]
	}
}
