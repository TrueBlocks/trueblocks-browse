package messages

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ErrorMsg struct {
	Address base.Address `json:"address"`
	ErrStr  string       `json:"errStr"`
}

func NewErrorMsg(err error, addrs ...base.Address) *ErrorMsg {
	addr := base.ZeroAddr
	if len(addrs) > 0 {
		addr = addrs[0]
	}

	return &ErrorMsg{
		Address: addr,
		ErrStr:  err.Error(),
	}
}
