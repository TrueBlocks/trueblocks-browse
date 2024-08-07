package daemons

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonFreshen struct {
	Daemon `json:"daemon"`
}

func NewFreshen(name string, sleep time.Duration) *DaemonFreshen {
	return &DaemonFreshen{
		Daemon: Daemon{
			Name:    name,
			Sleep:   sleep,
			Color:   "blue",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *DaemonFreshen) Run() {
	logger.Info("Starting fresheners...")

	for {
		if s.Daemon.State == Running {
			s.Daemon.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonFreshen) Stop() error {
	return s.Daemon.Stop()
}

func (s *DaemonFreshen) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonFreshen) Toggle() error {
	return s.Daemon.Toggle()
}

func (s *DaemonFreshen) Tick() int {
	return s.Daemon.Tick()
}
