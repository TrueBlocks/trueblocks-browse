package messages

import "context"

type InfoMsg struct {
	Message string `json:"message"`
}

func NewInfoMsg(message string) *InfoMsg {
	return &InfoMsg{Message: message}
}

func EmitInfo(ctx context.Context, msg string) {
	emitMsg(ctx, Info, NewInfoMsg(msg))
}

func (m *InfoMsg) Instance() InfoMsg {
	return InfoMsg{}
}
