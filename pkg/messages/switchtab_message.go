package messages

type SwitchTabMsg struct {
	Dest string `json:"dest"`
}

func NewSwitchTabMsg(message string) *SwitchTabMsg {
	return &SwitchTabMsg{Dest: message}
}

func (m *SwitchTabMsg) Instance() SwitchTabMsg {
	return SwitchTabMsg{}
}
