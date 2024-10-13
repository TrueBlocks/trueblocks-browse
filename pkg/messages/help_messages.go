package messages

type HelpMsg struct {
}

func NewHelpMsg() *HelpMsg {
	return &HelpMsg{}
}

func (m *HelpMsg) Instance() HelpMsg {
	return HelpMsg{}
}
