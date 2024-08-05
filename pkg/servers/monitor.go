package servers

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonFreshen struct {
	Server `json:"server"`
}

func NewFreshen(name string, sleep time.Duration) *DaemonFreshen {
	return &DaemonFreshen{
		Server: Server{
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
		if s.Server.State == Running {
			s.Server.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonFreshen) Stop() error {
	return s.Server.Stop()
}

func (s *DaemonFreshen) Pause() error {
	return s.Server.Pause()
}

func (s *DaemonFreshen) Toggle() error {
	return s.Server.Toggle()
}

func (s *DaemonFreshen) Tick() int {
	return s.Server.Tick()
}
