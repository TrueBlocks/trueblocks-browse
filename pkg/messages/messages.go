package messages

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message string

const (
	Completed Message = "Completed"
	Cancelled Message = "Cancelled"
	Error     Message = "Error"
	Warn      Message = "Warn"
	Info      Message = "Info"
	Progress  Message = "Progress"
	Daemon    Message = "Daemon"
	Document  Message = "Document"
	Navigate  Message = "Navigate"
	Reload    Message = "Reload"
	Help      Message = "Help"
)

var AllMessages = []struct {
	Value  Message `json:"value"`
	TSName string  `json:"tsname"`
}{
	{Completed, "COMPLETED"},
	{Cancelled, "CANCELLED"},
	{Error, "ERROR"},
	{Warn, "WARNING"},
	{Info, "INFO"},
	{Progress, "PROGRESS"},
	{Daemon, "DAEMON"},
	{Document, "DOCUMENT"},
	{Navigate, "NAVIGATE"},
	{Reload, "RELOAD"},
	{Help, "HELP"},
}

type MessageData interface {
	string | ProgressMsg | DaemonMsg | ErrorMsg | DocumentMsg | NavigateMsg | HelpMsg | InfoMsg | CancelMsg
}

func Send[T MessageData](ctx context.Context, msg Message, data *T) {
	runtime.EventsEmit(ctx, string(msg), data)
}
