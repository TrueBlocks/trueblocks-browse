package messages

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"

type ProgressMsg struct {
	Address base.Address `json:"address"`
	Have    int          `json:"have"`
	Want    int          `json:"want"`
}

func NewProgressMsg(address base.Address, have int, want int) *ProgressMsg {
	return &ProgressMsg{
		Address: address,
		Have:    have,
		Want:    want,
	}
}

func (m *ProgressMsg) Instance() ProgressMsg {
	return ProgressMsg{}
}
