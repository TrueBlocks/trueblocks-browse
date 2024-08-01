package servers

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type Ipfs struct {
	Server `json:"server"`
}

func NewIpfs(name string, sleep time.Duration) *Ipfs {
	return &Ipfs{
		Server: Server{
			Name:    name,
			Sleep:   sleep,
			Color:   "red",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *Ipfs) Run() {
	logger.Info("Starting ipfs...")

	for {
		if s.Server.State == Running {
			s.Server.Notify()
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *Ipfs) Stop() error {
	return s.Server.Stop()
}

func (s *Ipfs) Pause() error {
	return s.Server.Pause()
}

func (s *Ipfs) Toggle() error {
	return s.Server.Toggle()
}

func (s *Ipfs) Tick() int {
	return s.Server.Tick()
}
