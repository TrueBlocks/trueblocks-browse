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
			s.Tick("Tick")
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonIpfs) Tick(msg ...string) int {
	return s.Daemon.Tick(msg...)
}
