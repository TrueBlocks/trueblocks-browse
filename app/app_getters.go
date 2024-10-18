package app

import (
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) IsShowing(which string) bool {
	return a.session.Toggles[which]
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

// ------------------------------------------------------------------------
func (a *App) GetAddress() base.Address {
	addr := a.session.LastSub["/history"]
	return base.HexToAddress(addr)
}
