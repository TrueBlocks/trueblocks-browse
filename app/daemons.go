package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
)

func (a *App) StartDaemons() {
	go a.ScraperController.Run()
	go a.FileController.Run()
	go a.FreshenController.Run()
	go a.IpfsController.Run()
}

func (a *App) GetDaemon(name string) *daemons.Daemon {
	switch name {
	case "scraper":
		return &a.ScraperController.Daemon
	case "filedaemon":
		return &a.FileController.Daemon
	case "freshen":
		return &a.FreshenController.Daemon
	case "ipfs":
		return &a.IpfsController.Daemon
	default:
		return nil
	}
}

func (a *App) ToggleDaemon(name string) error {
	switch name {
	case "scraper":
		return a.ScraperController.Toggle()
	case "filedaemon":
		return a.FileController.Toggle()
	case "freshen":
		return a.FreshenController.Toggle()
	case "ipfs":
		return a.IpfsController.Toggle()
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
