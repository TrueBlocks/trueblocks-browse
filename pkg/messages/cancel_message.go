package messages

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"

type CancelMsg struct {
	Address base.Address `json:"address"`
}

func NewCancelMsg(addrs ...base.Address) *CancelMsg {
	address := base.ZeroAddr
	if len(addrs) > 0 {
		address = addrs[0]
	}

	return &CancelMsg{
		Address: address,
	}
}

func (m *CancelMsg) Instance() CancelMsg {
	return CancelMsg{}
}
