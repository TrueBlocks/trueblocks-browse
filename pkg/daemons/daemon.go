package daemons

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
)

type Freshener interface {
	Refresh(which ...string)
	GetContext() context.Context
}

type Daemon struct {
	Name      string        `json:"name"`
	Sleep     time.Duration `json:"sleep"`
	Color     string        `json:"color"`
	Started   time.Time     `json:"started"`
	Ticks     int           `json:"ticks"`
	state     State
	freshener Freshener
}

func (s *Daemon) Run() {
	s.state = Running
	s.Tick("Run")
}

func (s *Daemon) Stop() error {
	s.state = Stopped
	s.Tick("Stopped")
	return nil
}

func (s *Daemon) Pause() error {
	s.state = Paused
	s.Tick("Paused")
	return nil
}

func (s *Daemon) Toggle() error {
	if s.IsRunning() {
		return s.Pause()
	}
	s.Run()
	return nil
}

func (s *Daemon) Tick(msg ...string) int {
	msgOut := fmt.Sprintf("%-10.10s (% 5d-% 5.2f): %s",
		s.Name,
		s.Ticks,
		float64(time.Since(s.Started))/float64(time.Second),
		msg,
	)

	messages.Send(s.freshener.GetContext(), messages.Daemon, messages.NewDaemonMsg(
		strings.ToLower(s.Name),
		msgOut,
		s.Color,
	))
	s.Ticks++
	return s.Ticks
}

func (s *Daemon) IsRunning() bool {
	return s.state == Running
}

func (s *Daemon) StateToString() string {
	return s.state.String()
}

type Daemoner interface {
	Run()
	Stop() error
	Pause() error
	Tick(msg ...string) int
	IsRunning() bool
	// Toggle() error
}
