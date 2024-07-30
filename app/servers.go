package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/servers"
)

func (a *App) StartServers() {
	go a.Scraper.Run()
	go a.FileServer.Run()
	go a.Monitor.Run()
	go a.Ipfs.Run()
}

func (a *App) GetServer(name servers.Type) *servers.Server {
	switch name {
	case servers.ST_Scraper:
		return &a.Scraper.Server
	case servers.ST_FileServer:
		return &a.FileServer.Server
	case servers.ST_Monitor:
		return &a.Monitor.Server
	case servers.ST_Ipfs:
		return &a.Ipfs.Server
	default:
		return nil
	}
}

func (a *App) ToggleServer(name servers.Type) error {
	switch name {
	case servers.ST_Scraper:
		return a.Scraper.Server.Toggle()
	case servers.ST_FileServer:
		return a.FileServer.Toggle()
	case servers.ST_Monitor:
		return a.Monitor.Server.Toggle()
	case servers.ST_Ipfs:
		return a.Ipfs.Server.Toggle()
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
