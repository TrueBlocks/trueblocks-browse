package messages

type ToggleMsg struct {
	Component string `json:"component"`
}

func NewToggleMsg(comp string) *ToggleMsg {
	return &ToggleMsg{
		Component: comp,
	}
}

func (m *ToggleMsg) Instance() ToggleMsg {
	return ToggleMsg{}
}
