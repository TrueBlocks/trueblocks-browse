package messages

type NavigateMsg struct {
	Route string `json:"route"`
}

func NewNavigateMsg(route string) *NavigateMsg {
	return &NavigateMsg{
		Route: route,
	}
}

// This function is required for Wails to generate the binding code.
func (m *NavigateMsg) Instance() NavigateMsg {
	return NavigateMsg{}
}
