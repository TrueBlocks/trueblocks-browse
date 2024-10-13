package messages

type InfoMsg struct {
	Message string `json:"message"`
}

func NewInfoMsg(message string) *InfoMsg {
	return &InfoMsg{Message: message}
}

func (m *InfoMsg) Instance() InfoMsg {
	return InfoMsg{}
}
