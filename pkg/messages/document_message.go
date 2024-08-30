package messages

type DocumentMsg struct {
	Filename string `json:"filename"`
	Msg      string `json:"msg"`
}

func NewDocumentMsg(filename string, msg string) *DocumentMsg {
	return &DocumentMsg{
		Filename: filename,
		Msg:      msg,
	}
}

// This function is required for Wails to generate the binding code.
func (m *DocumentMsg) Instance() DocumentMsg {
	return DocumentMsg{}
}
