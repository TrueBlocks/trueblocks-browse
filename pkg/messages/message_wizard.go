package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
)

type WizardMsg struct {
	State wizard.State `json:"state"`
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
