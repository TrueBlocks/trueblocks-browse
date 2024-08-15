package messages

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message int

const (
	Completed Message = iota
	Error
	Warn
	Progress
	Daemon
	Document
)

type MessageData interface {
	string | ProgressMsg | DaemonMsg | ErrorMsg | DocumentMsg | NavigateMsg | HelpMsg
}

// AllMessages - all possible messages for the frontend codegen
var AllMessages = []struct {
	Value  Message `json:"value"`
	TSName string  `json:"tsname"`
}{
	{Completed, "COMPLETED"},
	{Error, "ERROR"},
	{Warn, "WARN"},
	{Progress, "PROGRESS"},
	{Daemon, "DAEMON"},
	{Document, "DOCUMENT"},
	{Navigate, "NAVIGATE"},
	{ToggleHelp, "TOGGLEHELP"},
}

func Send[T MessageData](ctx context.Context, msg Message, data *T) {
	runtime.EventsEmit(ctx, MessageType(msg), data)
}
