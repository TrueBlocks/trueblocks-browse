package messages

type NavigateMsg struct {
	Route string `json:"route"`
}

func NewNavigateMsg(route string) *NavigateMsg {
	return &NavigateMsg{
		Route: route,
	}
}

func (m *NavigateMsg) Instance() NavigateMsg {
	return NavigateMsg{}
}
