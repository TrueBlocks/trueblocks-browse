package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type InfoMsg struct {
	Message string `json:"message"`
}

func NewInfoMessage(message string) *InfoMsg {
	return &InfoMsg{Message: message}
}

func SendInfo(ctx context.Context, msg string) {
	logger.Info(msg)
	Send(ctx, Info, NewInfoMessage(msg))
}

// This function is required for Wails to generate the binding code.
func (m *InfoMsg) Instance() InfoMsg {
	return InfoMsg{}
}
