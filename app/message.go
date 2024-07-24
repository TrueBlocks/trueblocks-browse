package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message int

const (
	Completed Message = iota
	Error
	Warn
	Progress
)

type ProgressMsg struct {
	Address base.Address `json:"address"`
	Have    int64        `json:"have"`
	Want    int64        `json:"want"`
}

func (a *App) MessageType(msg Message) string {
	m := map[Message]string{
		Completed: "Completed",
		Error:     "Error",
		Warn:      "Warn",
		Progress:  "Progress",
	}
	return m[msg]
}

func (a *App) SendMessage(addr base.Address, msg Message, data interface{}) {
	runtime.EventsEmit(a.ctx, a.MessageType(msg), data)
}
