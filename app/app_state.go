package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) GetLast(which string) string {
	switch which {
	case "route":
		return a.GetSession().LastRoute + a.GetLastSub(a.GetSession().LastRoute)
	case "help":
		return a.GetSession().LastHelp
	}
	return "Unknown"
}

func (a *App) SetLast(which, value string) {
	switch which {
	case "route":
		parts := strings.Split(value, "/")
		if len(parts) > 2 {
			if !strings.HasPrefix(parts[2], ":") {
				route := "/" + parts[1]
				a.GetSession().LastRoute = route
				a.GetSession().LastSub[route] = parts[2]
			}
		} else {
			a.GetSession().LastRoute = value
		}
	case "help":
		a.GetSession().LastHelp = value
	}
	a.GetSession().Save()
}

func (a *App) GetLastSub(which string) string {
	val := a.GetSession().LastSub[which]
	if val == "" {
		return ""
	}
	return "/" + val
}

func (a *App) GetLastDaemon(which string) bool {
	switch which {
	case "daemon-freshen":
		return a.GetSession().Daemons.Freshen
	case "daemon-scraper":
		return a.GetSession().Daemons.Scraper
	case "daemon-ipfs":
		return a.GetSession().Daemons.Ipfs
	}
	logger.Error("Should not happen in GetLastDaemon")
	return false
}

func (a *App) SetLastDaemon(which string, value bool) {
	switch which {
	case "daemon-freshen":
		a.GetSession().Daemons.Freshen = value
	case "daemon-scraper":
		a.GetSession().Daemons.Scraper = value
	case "daemon-ipfs":
		a.GetSession().Daemons.Ipfs = value
	}
	a.GetSession().Save()
}
