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

func (m *DocumentMsg) Instance() DocumentMsg {
	return DocumentMsg{}
}
