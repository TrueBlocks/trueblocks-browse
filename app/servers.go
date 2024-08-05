package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/servers"
)

func (a *App) StartServers() {
	go a.ScraperController.Run()
	go a.FileController.Run()
	go a.FreshenController.Run()
	go a.IpfsController.Run()
}

func (a *App) GetServer(name string) *servers.Server {
	switch name {
	case "scraper":
		return &a.ScraperController.Server
	case "fileserver":
		return &a.FileController.Server
	case "freshen":
		return &a.FreshenController.Server
	case "ipfs":
		return &a.IpfsController.Server
	default:
		return nil
	}
}

func (a *App) ToggleServer(name string) error {
	switch name {
	case "scraper":
		return a.ScraperController.Toggle()
	case "fileserver":
		return a.FileController.Toggle()
	case "freshen":
		return a.FreshenController.Toggle()
	case "ipfs":
		return a.IpfsController.Toggle()
	default:
		return fmt.Errorf("server %s not found in ToggleServer", name)
	}
}

func (a *App) StateToString(name string) string {
	s := a.GetServer(name)
	if s == nil {
		return "Server not found"
	}
	return s.State.String()
}
