package daemons

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonIpfs struct {
	Daemon `json:"daemon"`
}

func NewIpfs(freshener Freshener, name string, sleep time.Duration) *DaemonIpfs {
	return &DaemonIpfs{
		Daemon: Daemon{
			Name:      name,
			Sleep:     sleep,
			Color:     "red",
			State:     Paused,
			Started:   time.Now(),
			freshener: freshener,
		},
	}
}

func (s *DaemonIpfs) Run() {
	logger.Info("Starting ipfs...")

	for {
		if s.Daemon.State == Running {
			s.Daemon.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonIpfs) Stop() error {
	return s.Daemon.Stop()
}

func (s *DaemonIpfs) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonIpfs) Toggle() error {
	return s.Daemon.Toggle()
}

func (s *DaemonIpfs) Tick() int {
	return s.Daemon.Tick()
}
