package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

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
	a.saveSession()
}
