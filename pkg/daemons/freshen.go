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
			Started:   time.Now(),
			State:     state,
			freshener: freshener,
		},
	}
}

func (s *DaemonFreshen) String() string {
	return s.Daemon.String()
}

func (s *DaemonFreshen) GetState() DaemonState {
	return s.Daemon.GetState()
}

func (s *DaemonFreshen) IsRunning() bool {
	return s.Daemon.IsRunning()
}

func (s *DaemonFreshen) Run() {
	logger.Info("Starting fresheners...")
	for {
		if s.IsRunning() {
			s.Tick("Freshen")
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonFreshen) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonFreshen) Tick(msg ...string) int {
	s.freshener.Freshen()
	s.Ticks++
	return s.Ticks // we don't use the Daemon's Tick since Freshen notifies if it runs
}

func (s *DaemonFreshen) Toggle() error {
	return s.Daemon.Toggle()
}
