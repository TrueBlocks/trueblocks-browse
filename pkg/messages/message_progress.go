package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ProgressMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewProgressMsg(address base.Address, have, want int) *ProgressMsg {
	return &ProgressMsg{
		Address: address,
		Int1:    have,
		Int2:    want,
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
