package servers

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type Monitor struct {
	Server `json:"server"`
}

func NewMonitor(name string, sleep time.Duration) *Monitor {
	return &Monitor{
		Server: Server{
			Name:    name,
			Sleep:   sleep,
			Color:   "blue",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *Monitor) Run() {
	logger.Info("Starting monitors...")

	for {
		if s.Server.State == Running {
			s.Server.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *Monitor) Stop() error {
	return s.Server.Stop()
}

func (s *Monitor) Pause() error {
	return s.Server.Pause()
}

func (s *Monitor) Toggle() error {
	return s.Server.Toggle()
}

func (s *Monitor) Tick() int {
	return s.Server.Tick()
}
