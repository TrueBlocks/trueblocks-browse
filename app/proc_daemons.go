package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
)

func (a *App) startDaemons() {
	go a.FreshenController.Run()
	go a.ScraperController.Run()
	go a.IpfsController.Run()
}

func (a *App) GetDaemonJson(name string) string {
	d := a.getDaemon(name)
	if d == nil {
		return "{}"
	}
	return d.String()
}

func (a *App) ToggleDaemon(name string) error {
	switch name {
	case "freshen":
		err := a.FreshenController.Toggle()
		a.SetLastDaemon("daemon-freshen", a.FreshenController.IsRunning())
		return err
	case "scraper":
		err := a.ScraperController.Toggle()
		a.SetLastDaemon("daemon-scraper", a.ScraperController.IsRunning())
		return err
	case "ipfs":
		err := a.IpfsController.Toggle()
		a.SetLastDaemon("daemon-ipfs", a.IpfsController.IsRunning())
		return err
	default:
		return fmt.Errorf("daemon %s not found in ToggleDaemon", name)
	}
}

func (a *App) StateToString(name string) string {
	if s := a.getDaemon(name); s == nil {
		return "Daemon not found"
	} else {
		return s.StateToString()
	}
}

func (a *App) DaemonInstance() *daemons.Daemon {
	return &daemons.Daemon{}
}

func (a *App) getDaemon(name string) *daemons.Daemon {
	switch name {
	case "freshen":
		return &a.FreshenController.Daemon
	case "scraper":
		return &a.ScraperController.Daemon
	case "ipfs":
		return &a.IpfsController.Daemon
	default:
		return nil
	}
}
