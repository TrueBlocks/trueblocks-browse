package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) startDaemons() {
	go a.FreshenController.Run()
	go a.ScraperController.Run()
	go a.IpfsController.Run()
}

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
		return a.FreshenController
	case "scraper":
		return a.ScraperController
	case "ipfs":
		return a.IpfsController
	default:
		if len(name) > 0 {
			logger.Fatal("getDaemon", "should not happen", name)
		}
		return nil
	}
}
