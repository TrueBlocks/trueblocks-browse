package app

import (
	"context"
	"os"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) IsShowing(which string) bool {
	return a.session.Toggles.IsOn(which)
}

func (a *App) GetConfig() *coreConfig.ConfigFile {
	return &a.cfg
}

func (a *App) GetSession() *config.Session {
	return &a.session
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) GetWindow() *config.Window {
	return &a.session.Window
}

func (a *App) GetEnv(key string) string {
	return os.Getenv(key)
}

func (a *App) GetMeta() coreTypes.MetaData {
	return a.meta
}

func (a *App) GetAppTitle() string {
	return a.session.Window.Title
}

func (a *App) GetRoute() string {
	if !a.IsConfigured() {
		return "/wizard"
	}

	route := a.session.LastRoute
	if len(a.session.LastSub) > 0 {
		route += "/" + a.session.LastSub[route]
	}

	return route
}

func (a *App) GetAddress() base.Address {
	addr := a.session.LastSub["/history"]
	return base.HexToAddress(addr)
}

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
