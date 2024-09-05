package messages

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

// This function is required for Wails to generate the binding code.
func (m *DaemonMsg) Instance() DaemonMsg {
	return DaemonMsg{}
}
