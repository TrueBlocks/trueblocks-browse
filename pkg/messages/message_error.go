package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ErrorMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   wizard.State `json:"state"`
	Int1    int          `json:"int1"`
	Int2    int          `json:"int2"`
	Msg1    string       `json:"msg1"`
	Msg2    string       `json:"msg2"`
}

func NewErrorMsg(err error, addrs ...base.Address) *ErrorMsg {
	address := base.ZeroAddr
	if len(addrs) > 0 {
		address = addrs[0]
	}

	return &ErrorMsg{
		Address: address,
		Msg1:    err.Error(),
	}
}

func EmitError(ctx context.Context, err error, addrs ...base.Address) {
	emitMsg(ctx, Error, NewErrorMsg(err, addrs...))
}

func (m *ErrorMsg) Instance() ErrorMsg {
	return ErrorMsg{}
}
