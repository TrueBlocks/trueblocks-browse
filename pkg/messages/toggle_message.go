package messages

type ToggleMsg struct {
	Layout string `json:"layout"`
}

func NewToggleMsg(comp string) *ToggleMsg {
	return &ToggleMsg{
		Layout: comp,
	}
}

func (m *ToggleMsg) Instance() ToggleMsg {
	return ToggleMsg{}
}
