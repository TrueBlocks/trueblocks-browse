package servers

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonIpfs struct {
	Server `json:"server"`
}

func NewIpfs(name string, sleep time.Duration) *DaemonIpfs {
	return &DaemonIpfs{
		Server: Server{
			Name:    name,
			Sleep:   sleep,
			Color:   "red",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *DaemonIpfs) Run() {
	logger.Info("Starting ipfs...")

	for {
		if s.Server.State == Running {
			s.Server.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonIpfs) Stop() error {
	return s.Server.Stop()
}

func (s *DaemonIpfs) Pause() error {
	return s.Server.Pause()
}

func (s *DaemonIpfs) Toggle() error {
	return s.Server.Toggle()
}

func (s *DaemonIpfs) Tick() int {
	return s.Server.Tick()
}
