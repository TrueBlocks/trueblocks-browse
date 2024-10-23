package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type WizardMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewWizardMsg(state wizard.State) *WizardMsg {
	return &WizardMsg{
		State: state,
	}
}

func EmitWizard(ctx context.Context, state wizard.State) {
	emitMsg(ctx, Wizard, NewWizardMsg(state))
}

func (m *WizardMsg) Instance() WizardMsg {
	return WizardMsg{}
}
