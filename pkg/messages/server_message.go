package messages

type ServerMsg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Color   string `json:"color"`
}

func NewServerMsg(name string, message string, color string) *ServerMsg {
	return &ServerMsg{
		Name:    name,
		Message: message,
		Color:   color,
	}
}
