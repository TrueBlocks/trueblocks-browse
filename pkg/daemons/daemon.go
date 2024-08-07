package daemons

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
)

type Freshener interface {
	// DaemonMsg(msg *DaemonMsg)
	GetContext() context.Context
}

type Daemon struct {
	Name      string        `json:"name"`
	Sleep     time.Duration `json:"sleep"`
	Color     string        `json:"color"`
	Started   time.Time     `json:"started"`
	Runs      int           `json:"runs"`
	State     State         `json:"state"`
	freshener Freshener
}

func (s *Daemon) Run() error {
	s.State = Running
	s.Notify("Run")
	return nil
}

func (s *Daemon) Stop() error {
	s.State = Stopped
	s.Notify("Stopped")
	return nil
}

func (s *Daemon) Pause() error {
	s.State = Paused
	s.Notify("Paused")
	return nil
}

func (s *Daemon) Toggle() error {
	if s.State == Running {
		return s.Pause()
	}
	return s.Run()
}

func (s *Daemon) Tick() int {
	s.Runs++
	return s.Runs
}

func (s *Daemon) Notify(msg ...string) {
	color := colors.ColorMap[s.Color]
	if color == "" {
		color = colors.White
	}

	s.Tick()
	msgOut := fmt.Sprintf("%-10.10s (% 5d-% 5.2f): %s",
		s.Name,
		s.Runs,
		float64(time.Since(s.Started))/float64(time.Second),
		msg,
	)
	messages.Send(s.messenger.GetContext(), messages.Daemon, messages.NewDaemonMsg(
		strings.ToLower(s.Name),
		msgOut,
		color,
	))
}

type Daemoner *interface {
	Run() error
	Stop() error
	Pause() error
	Tick() int
}
