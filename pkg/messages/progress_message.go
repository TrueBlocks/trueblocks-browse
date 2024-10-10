package messages

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"

type ProgressMsg struct {
	Address base.Address `json:"address"`
	Have    int          `json:"have"`
	Want    int          `json:"want"`
}

func NewCancelMsg(addrs ...base.Address) *ProgressMsg {
	return NewProgressMsg(-1, -1, addrs...)
}

func NewProgressMsg(have int, want int, addrs ...base.Address) *ProgressMsg {
	address := base.ZeroAddr
	if len(addrs) > 0 {
		address = addrs[0]
	}

	return &ProgressMsg{
		Address: address,
		Have:    have,
		Want:    want,
	}
}

// This function is required for Wails to generate the binding code.
func (m *ProgressMsg) Instance() ProgressMsg {
	return ProgressMsg{}
}
