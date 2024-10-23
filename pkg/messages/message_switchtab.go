package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type SwitchTabMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewSwitchTabMsg(direction string) *SwitchTabMsg {
	return &SwitchTabMsg{Msg1: direction}
}

func EmitSwitchTab(ctx context.Context, msg string) {
	emitMsg(ctx, SwitchTab, NewSwitchTabMsg(msg))
}

func (m *SwitchTabMsg) Instance() SwitchTabMsg {
	return SwitchTabMsg{}
}
