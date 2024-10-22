package messages

import "context"

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

func EmitDocument(ctx context.Context, fileName, msg string) {
	emitMsg(ctx, Document, NewDocumentMsg(fileName, msg))
}

func (m *DocumentMsg) Instance() DocumentMsg {
	return DocumentMsg{}
}
