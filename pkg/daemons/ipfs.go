package daemons

import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonIpfs struct {
	Daemon `json:"daemon"`
}

func NewIpfs(freshener Freshener, name string, sleep time.Duration, start bool) *DaemonIpfs {
	state := Paused
	if start {
		state = Running
	}
	return &DaemonIpfs{
		Daemon: Daemon{
			Name:      name,
			Sleep:     sleep,
			Color:     "red",
			Started:   time.Now(),
			State:     state,
			freshener: freshener,
		},
	}
}

func (s *DaemonIpfs) String() string {
	bytes, _ := json.Marshal(s.Daemon)
	return string(bytes)
}

func (s *DaemonIpfs) Run() {
	logger.Info("Starting ipfs...")

	for {
		if s.IsRunning() {
			s.Tick("Tick")
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonIpfs) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonIpfs) Tick(msg ...string) int {
	return s.Daemon.Tick(msg...)
}

func (s *DaemonIpfs) IsRunning() bool {
	return s.Daemon.IsRunning()
}
