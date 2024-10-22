package messages

import "context"

type DaemonMsg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Color   string `json:"color"`
}

func NewDaemonMsg(name string, message string, color string) *DaemonMsg {
	return &DaemonMsg{
		Name:    name,
		Message: message,
		Color:   color,
	}
}

func EmitDaemon(ctx context.Context, name, msg, color string) {
	emitMsg(ctx, Daemon, NewDaemonMsg(name, msg, color))
}

func (m *DaemonMsg) Instance() DaemonMsg {
	return DaemonMsg{}
}
