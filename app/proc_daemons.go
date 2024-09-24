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

func (a *App) GetDaemon(name string) *daemons.Daemon {
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

func (a *App) ToggleDaemon(name string) error {
	switch name {
	case "freshen":
		err := a.FreshenController.Toggle()
		a.SetLastDaemon("daemon-freshen", a.FreshenController.State == daemons.Running)
		return err
	case "scraper":
		err := a.ScraperController.Toggle()
		a.SetLastDaemon("daemon-scraper", a.ScraperController.State == daemons.Running)
		return err
	case "ipfs":
		err := a.IpfsController.Toggle()
		a.SetLastDaemon("daemon-ipfs", a.IpfsController.State == daemons.Running)
		return err
	default:
		return fmt.Errorf("daemon %s not found in ToggleDaemon", name)
	}
}

func (a *App) StateToString(name string) string {
	s := a.GetDaemon(name)
	if s == nil {
		return "Daemon not found"
	}
	return s.State.String()
}
