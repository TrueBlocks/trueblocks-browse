package messages

import "context"

type SwitchTabMsg struct {
	Dest string `json:"dest"`
}

func NewSwitchTabMsg(message string) *SwitchTabMsg {
	return &SwitchTabMsg{Dest: message}
}

func EmitSwitchTab(ctx context.Context, msg string) {
	emitMsg(ctx, SwitchTab, NewSwitchTabMsg(msg))
}

func (m *SwitchTabMsg) Instance() SwitchTabMsg {
	return SwitchTabMsg{}
}
