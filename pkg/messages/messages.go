package messages

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message string

const (
	Completed  Message = "Completed"
	Error      Message = "Error"
	Warn       Message = "Warn"
	Progress   Message = "Progress"
	Daemon     Message = "Daemon"
	Document   Message = "Document"
	Navigate   Message = "Navigate"
	Reload     Message = "Reload"
	ToggleHelp Message = "ToggleHelp"
)

type MessageData interface {
	string | ProgressMsg | DaemonMsg | ErrorMsg | DocumentMsg | NavigateMsg | ReloadMsg | HelpMsg
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
	{Reload, "RELOAD"},
	{ToggleHelp, "TOGGLEHELP"},
}

func Send[T MessageData](ctx context.Context, msg Message, data *T) {
	runtime.EventsEmit(ctx, string(msg), data)
}
