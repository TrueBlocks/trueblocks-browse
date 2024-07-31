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
	Server
	Document
)

type MessageData interface {
	string | ProgressMsg | ServerMsg | ErrorMsg | DocumentMsg
}

var Messages = []struct {
	Value  Message
	TSName string
}{
	{Completed, "COMPLETED"},
	{Error, "ERROR"},
	{Warn, "WARN"},
	{Progress, "PROGRESS"},
	{Server, "SERVER"},
	{Document, "DOCUMENT"},
}

func MessageType(msg Message) string {
	m := make(map[Message]string, len(Messages))
	for _, message := range Messages {
		m[message.Value] = message.TSName
	}
	return m[msg]
}

func Send[T MessageData](ctx context.Context, msg Message, data *T) {
	runtime.EventsEmit(ctx, MessageType(msg), data)
}
