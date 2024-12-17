package app

import (
	"context"
	"os"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

func (a *App) GetConfig() *types.Config {
	return &a.config.Config
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) GetEnv(key string) string {
	return os.Getenv(key)
}

func (a *App) GetChains() []string {
	ret := []string{}
	if len(a.GetConfig().Chains) == 0 {
		ret = append(ret, "mainnet")
	} else {
		for _, chain := range a.GetConfig().Chains {
			ret = append(ret, chain.Chain)
		}
	}
	sort.Strings(ret)
	return ret
}

func (a *App) GetChainInfo(chain string) types.Chain {
	for _, ch := range a.status.Chains {
		if ch.Chain == chain {
			return ch
		}
	}
	return types.Chain{}
}
