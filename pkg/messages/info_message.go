package messages

import "context"

type InfoMsg struct {
	Message string `json:"message"`
}

func NewInfoMessage(message string) *InfoMsg {
	return &InfoMsg{Message: message}
}

func SendInfo(ctx context.Context, msg string) {
	Send(ctx, Info, NewInfoMessage(msg))
}

// This function is required for Wails to generate the binding code.
func (m *InfoMsg) Instance() InfoMsg {
	return InfoMsg{}
}
