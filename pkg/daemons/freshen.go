package daemons

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonFreshen struct {
	Daemon `json:"daemon"`
}

func NewFreshen(freshener Freshener, name string, sleep time.Duration, start bool) *DaemonFreshen {
	state := Paused
	if start {
		state = Running
	}
	return &DaemonFreshen{
		Daemon: Daemon{
			Name:      name,
			Sleep:     sleep,
			Color:     "blue",
			State:     state,
			Started:   time.Now(),
			freshener: freshener,
		},
	}
}

func (s *DaemonFreshen) Run() {
	logger.Info("Starting fresheners...")
	for {
		if s.Daemon.State == Running {
			s.Tick("Freshen")
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonFreshen) Tick(msg ...string) int {
	go s.freshener.Refresh()
	s.Ticks++
	return s.Ticks // we don't use the Daemon's Tick since Freshen notifies if it runs
}
