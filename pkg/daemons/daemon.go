package daemons

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
)

type Freshener interface {
	Refresh(skipable bool, which ...string)
	GetContext() context.Context
}

type Daemoner interface {
	String() string
	GetState() State
	IsRunning() bool
	Run()
	Stop() error
	Pause() error
	Tick(msg ...string) int
	Toggle() error
}

type Daemon struct {
	Name      string        `json:"name"`
	Sleep     time.Duration `json:"sleep"`
	Color     string        `json:"color"`
	Started   time.Time     `json:"started"`
	Ticks     int           `json:"ticks"`
	State     State         `json:"state"`
	freshener Freshener
}

func (s *Daemon) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Daemon) GetState() State {
	return s.State
}

func (s *Daemon) IsRunning() bool {
	return s.State == Running
}

func (s *Daemon) Run() {
	s.State = Running
	s.Tick("Run")
}

func (s *Daemon) Stop() error {
	s.State = Stopped
	s.Tick("Stopped")
	return nil
}

func (s *Daemon) Pause() error {
	s.State = Paused
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

func (s *Daemon) Instance() *Daemon {
	return &Daemon{}
}
