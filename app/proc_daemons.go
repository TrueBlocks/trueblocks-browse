package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
)

func (a *App) ToggleDaemon(name string) error {
	d := a.getDaemon(name)
	if err := d.Toggle(); err != nil {
		return err
	}
	a.SetShowing(name, d.IsRunning())
	return nil
}

func (a *App) GetDaemon(name string) string {
	return a.getDaemon(name).String()
}

func (a *App) GetState(name string) string {
	return a.getDaemon(name).GetState().String()
}

func (a *App) getDaemon(name string) daemons.Daemoner {
	switch name {
	case "freshen":
		return a.daemons.FreshenController
	case "scraper":
		return a.daemons.ScraperController
	case "ipfs":
		return a.daemons.IpfsController
	default:
		// logger.Fatal("getDaemon", "should not happen", name)
		return &daemons.Daemon{}
	}
}
