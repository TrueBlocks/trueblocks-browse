package messages

type HelpMsg struct {
}

func NewHelpMsg() *HelpMsg {
	return &HelpMsg{}
}

// This function is required for Wails to generate the binding code.
func (m *HelpMsg) Instance() HelpMsg {
	return HelpMsg{}
}
