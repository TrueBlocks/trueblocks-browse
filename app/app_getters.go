package app

import (
	"context"
	"os"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetConfig() *types.Config {
	return &a.config.Config
}

func (a *App) GetSession() *types.Session {
	return &a.session.Session
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) GetWindow() *types.Window {
	return &a.session.Window
}

func (a *App) GetEnv(key string) string {
	return os.Getenv(key)
}

func (a *App) GetAppTitle() string {
	return a.session.Window.Title
}

func (a *App) GetRoute() string {
	if !a.isConfigured() {
		return "/wizard"
	}

	route := a.session.LastRoute
	sub, _ := a.session.LastSub.Load(route)
	if len(sub) > 0 {
		route += "/" + sub
	}

	return route
}

func (a *App) GetSelected() base.Address {
	addr, _ := a.session.LastSub.Load("/history")
	return base.HexToAddress(addr)
}

func (a *App) getChain() string {
	return a.session.LastChain
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
