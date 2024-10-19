package messages

type ToggleMsg struct {
	Layout string `json:"layout"`
	Route  string `json:"route"`
}

func NewToggleMsg(comp, route string) *ToggleMsg {
	return &ToggleMsg{
		Layout: comp,
		Route:  route,
	}
}

func (m *ToggleMsg) Instance() ToggleMsg {
	return ToggleMsg{}
}
