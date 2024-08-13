package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) GetLast(which string) string {
	switch which {
	case "route":
		return a.GetSession().LastRoute
	case "address":
		return a.GetSession().LastAddress
	case "help":
		return a.GetSession().LastHelp
	}
	return "Unknown"
}

func (a *App) SetLast(which, value string) {
	switch which {
	case "route":
		a.GetSession().LastRoute = value
	case "address":
		a.GetSession().LastAddress = value
	case "help":
		a.GetSession().LastHelp = value
	}
	a.GetSession().Save()
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
