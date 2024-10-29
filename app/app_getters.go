package app

import (
	"context"
	"os"
	"sort"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) IsShowing(which string) bool {
	return a.sessions.Toggles.IsOn(which)
}

func (a *App) GetConfig() *configTypes.Config {
	return &a.cfg
}

func (a *App) GetSession() *coreTypes.Session {
	return &a.sessions.Session
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) GetWindow() *coreTypes.Window {
	return &a.sessions.Window
}

func (a *App) GetEnv(key string) string {
	return os.Getenv(key)
}

func (a *App) GetMeta() coreTypes.MetaData {
	return a.meta
}

func (a *App) GetAppTitle() string {
	return a.sessions.Window.Title
}

func (a *App) GetRoute() string {
	if !a.IsConfigured() {
		return "/wizard"
	}

	route := a.sessions.LastRoute
	if len(a.sessions.LastSub) > 0 {
		route += "/" + a.sessions.LastSub[route]
	}

	return route
}

func (a *App) GetAddress() base.Address {
	addr := a.sessions.LastSub["/history"]
	return base.HexToAddress(addr)
}

func (a *App) GetChain() string {
	return a.Chain
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
