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
	return a.session.Toggles.IsOn(which)
}

func (a *App) GetConfig() *configTypes.Config {
	return &a.config.Config
}

func (a *App) GetSession() *coreTypes.Session {
	return &a.session.Session
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) GetWindow() *coreTypes.Window {
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
	if len(a.session.LastSub) > 0 {
		route += "/" + a.session.LastSub[route]
	}

	return route
}

func (a *App) GetSelected() base.Address {
	addr := a.session.LastSub["/history"]
	return base.HexToAddress(addr)
}

func (a *App) getChain() string {
	return a.Chain
}

func (a *App) GetChains() []string {
	ret := []string{}
	for _, chain := range a.GetConfig().Chains {
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
