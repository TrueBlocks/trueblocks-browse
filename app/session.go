package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// ------------------------------------------------------------------------
func (a *App) SetSessionVal(which, value string) {
	switch which {
	case "file":
		a.session.LastFile = value
	case "route":
		parts := strings.Split(value, "/")
		if len(parts) > 2 {
			if !strings.HasPrefix(parts[2], ":") {
				route := "/" + parts[1]
				a.session.LastRoute = route
				a.session.LastSub[route] = parts[2]
			}
		} else {
			a.session.LastRoute = value
		}
	case "chain":
		a.session.Chain = value
	}
	a.session.Save()
}

// ------------------------------------------------------------------------
func (a *App) GetSessionSubVal(which string) string {
	val := a.session.LastSub[which]
	if val == "" {
		return ""
	}
	return "/" + val
}

// ------------------------------------------------------------------------
func (a *App) GetSessionDeamon(which string) bool {
	switch which {
	case "daemon-freshen":
		return a.session.Daemons.Freshen
	case "daemon-scraper":
		return a.session.Daemons.Scraper
	case "daemon-ipfs":
		return a.session.Daemons.Ipfs
	}
	logger.Error("Should not happen in GetSessionDeamon")
	return false
}

// ------------------------------------------------------------------------
func (a *App) SetSessionDaemon(which string, value bool) {
	switch which {
	case "daemon-freshen":
		a.session.Daemons.Freshen = value
	case "daemon-scraper":
		a.session.Daemons.Scraper = value
	case "daemon-ipfs":
		a.session.Daemons.Ipfs = value
	}
	a.session.Save()
}

// ------------------------------------------------------------------------
func (a *App) GetLastAddress() base.Address {
	val := a.GetSessionSubVal("/history")
	return base.HexToAddress(strings.ReplaceAll(val, "/", ""))
}
