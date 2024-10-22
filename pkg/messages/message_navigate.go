package messages

import "context"

type NavigateMsg struct {
	Route string `json:"route"`
}

func NewNavigateMsg(route string) *NavigateMsg {
	return &NavigateMsg{
		Route: route,
	}
}

func EmitNavigate(ctx context.Context, route string) {
	emitMsg(ctx, Navigate, NewNavigateMsg(route))
}

func (m *NavigateMsg) Instance() NavigateMsg {
	return NavigateMsg{}
}
