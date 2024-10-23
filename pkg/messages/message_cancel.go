package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func NewCancelMsg(addrs ...base.Address) *MessageMsg {
	address := base.ZeroAddr
	if len(addrs) > 0 {
		address = addrs[0]
	}

	return &MessageMsg{
		Address: address,
	}
}

func EmitCancel(ctx context.Context, addrs ...base.Address) {
	emitMsg(ctx, Cancelled, NewCancelMsg(addrs...))
}
