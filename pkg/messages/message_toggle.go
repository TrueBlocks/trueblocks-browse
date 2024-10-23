package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ToggleMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewToggleMsg(comp, route string) *ToggleMsg {
	return &ToggleMsg{
		Msg1: route,
		Msg2: comp,
	}
}

func EmitToggle(ctx context.Context, comp, route string) {
	if comp != "" {
		emitMsg(ctx, ToggleLayout, NewToggleMsg(comp, ""))
	}
	// do not collapse
	if route != "" {
		emitMsg(ctx, ToggleHeader, NewToggleMsg("", route))
	}
}

func (m *ToggleMsg) Instance() ToggleMsg {
	return ToggleMsg{}
}
