package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ErrorMsg struct {
	Address base.Address `json:"address"`
	ErrStr  string       `json:"errStr"`
}

func NewErrorMsg(err error, addrs ...base.Address) *ErrorMsg {
	address := base.ZeroAddr
	if len(addrs) > 0 {
		address = addrs[0]
	}

	return &ErrorMsg{
		Address: address,
		ErrStr:  err.Error(),
	}
}

func EmitError(ctx context.Context, err error, addrs ...base.Address) {
	emitMsg(ctx, Error, NewErrorMsg(err, addrs...))
}

func (m *ErrorMsg) Instance() ErrorMsg {
	return ErrorMsg{}
}
