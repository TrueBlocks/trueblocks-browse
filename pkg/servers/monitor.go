package servers

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonMonitor struct {
	Server `json:"server"`
}

func NewMonitor(name string, sleep time.Duration) *DaemonMonitor {
	return &DaemonMonitor{
		Server: Server{
			Name:    name,
			Sleep:   sleep,
			Color:   "blue",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *DaemonMonitor) Run() {
	logger.Info("Starting monitors...")

	for {
		if s.Server.State == Running {
			s.Server.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonMonitor) Stop() error {
	return s.Server.Stop()
}

func (s *DaemonMonitor) Pause() error {
	return s.Server.Pause()
}

func (s *DaemonMonitor) Toggle() error {
	return s.Server.Toggle()
}

func (s *DaemonMonitor) Tick() int {
	return s.Server.Tick()
}
