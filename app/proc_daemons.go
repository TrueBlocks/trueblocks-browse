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
	if s := a.getDaemon(name); s == nil {
		return fmt.Errorf("could not find daemon %s", name)
	} else {
		if err := s.Toggle(); err != nil {
			return err
		}
		a.SetLastDaemon("daemon-"+name, s.IsRunning())
		return nil
	}
}

func (a *App) StateToString(name string) string {
	if s := a.getDaemon2(name); s == nil {
		return "Daemon not found"
	} else {
		return s.GetState().String()
	}
}

func (a *App) getDaemon2(name string) daemons.Daemoner {
	switch name {
	case "freshen":
		return a.FreshenController
	case "scraper":
		return a.ScraperController
	case "ipfs":
		return a.IpfsController
	default:
		return nil
	}
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
