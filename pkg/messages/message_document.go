package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type DocumentMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewDocumentMsg(filename string, operation string) *DocumentMsg {
	return &DocumentMsg{
		Msg1: filename,
		Msg2: operation,
	}
}

func EmitDocument(ctx context.Context, fileName, operation string) {
	emitMsg(ctx, Document, NewDocumentMsg(fileName, operation))
}

func (m *DocumentMsg) Instance() DocumentMsg {
	return DocumentMsg{}
}
