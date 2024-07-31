package servers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
)

type Server struct {
	Name    string          `json:"name"`
	Sleep   time.Duration   `json:"sleep"`
	Color   string          `json:"color"`
	Started time.Time       `json:"started"`
	Runs    int             `json:"runs"`
	State   State           `json:"state"`
	MsgCtx  context.Context `json:"-"`
}

func (s *Server) Run() error {
	s.State = Running
	s.Notify("Run")
	return nil
}

func (s *Server) Stop() error {
	s.State = Stopped
	s.Notify("Stopped")
	return nil
}

func (s *Server) Pause() error {
	s.State = Paused
	s.Notify("Paused")
	return nil
}

func (s *Server) Toggle() error {
	if s.State == Running {
		return s.Pause()
	}
	return s.Run()
}

func (s *Server) Tick() int {
	s.Runs++
	return s.Runs
}

func (s *Server) Notify(msg ...string) {
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
	// fmt.Printf("%sNotify: %s%s\n", color, msgOut, colors.Off)
	messages.Send(s.MsgCtx, messages.Server, messages.NewServerMsg(
		strings.ToLower(s.Name),
		msgOut,
		color,
	))
}

type Serverer *interface {
	Run() error
	Stop() error
	Pause() error
	Tick() int
}
