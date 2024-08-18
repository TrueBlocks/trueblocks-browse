package app

import (
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) getBalance(address base.Address) string {
	b, exists := a.balanceMap.Load(address)
	if exists {
		return b.(string)
	}

	opts := sdk.StateOptions{
		Addrs: []string{address.Hex()},
		Globals: sdk.Globals{
			Ether: true,
			Cache: true,
		},
	}
	if balances, meta, err := opts.State(); err != nil {
		return "0"
	} else {
		a.meta = *meta
		value := balances[0].Balance.ToEtherStr(18)
		a.balanceMap.Store(address, value)
		return value
	}
}
