package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ProgressMsg struct {
	Address base.Address `json:"address"`
	Have    int          `json:"have"`
	Want    int          `json:"want"`
}

func NewProgressMsg(address base.Address, have int, want int) *ProgressMsg {
	return &ProgressMsg{
		Address: address,
		Have:    have,
		Want:    want,
	}
}

func EmitCompleted(ctx context.Context, address base.Address, total int) {
	emitMsg(ctx, Completed, NewProgressMsg(address, total, total))
}

func EmitProgress(ctx context.Context, address base.Address, have, want int) {
	emitMsg(ctx, Progress, NewProgressMsg(address, have, want))

}

func (m *ProgressMsg) Instance() ProgressMsg {
	return ProgressMsg{}
}
