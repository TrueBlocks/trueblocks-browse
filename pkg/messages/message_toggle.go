package messages

import "context"

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

func EmitToggle(ctx context.Context, comp, route string) {
	if comp != "" {
		emitMsg(ctx, ToggleLayout, NewToggleMsg(comp, ""))
	}
	// do not collapse
	if route != "" {
		emitMsg(ctx, ToggleHeader, NewToggleMsg("", route))
	}
}

func (m *ToggleMsg) Instance() ToggleMsg {
	return ToggleMsg{}
}
