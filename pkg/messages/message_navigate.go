package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type NavigateMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewNavigateMsg(route string) *NavigateMsg {
	return &NavigateMsg{
		Msg1: route,
	}
}

func EmitNavigate(ctx context.Context, route string) {
	emitMsg(ctx, Navigate, NewNavigateMsg(route))
}

func (m *NavigateMsg) Instance() NavigateMsg {
	return NavigateMsg{}
}
