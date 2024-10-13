package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type InfoMsg struct {
	Message string `json:"message"`
}

func NewInfoMsg(message string) *InfoMsg {
	return &InfoMsg{Message: message}
}

func SendInfo(ctx context.Context, msg string) {
	logger.Info(msg)
	Send(ctx, Info, NewInfoMsg(msg))
}

func (m *InfoMsg) Instance() InfoMsg {
	return InfoMsg{}
}
