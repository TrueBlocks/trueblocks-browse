package messages

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"

type ProgressMsg struct {
	Address base.Address `json:"address"`
	Have    int64        `json:"have"`
	Want    int64        `json:"want"`
}

func NewProgressMsg(have int64, want int64, addrs ...base.Address) *ProgressMsg {
	addr := base.ZeroAddr
	if len(addrs) > 0 {
		addr = addrs[0]
	}

	return &ProgressMsg{
		Address: addr,
		Have:    have,
		Want:    want,
	}
}
