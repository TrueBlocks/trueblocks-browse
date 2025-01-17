package messages

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type MessageMsg struct {
	Name    string       `json:"name"`
	Address base.Address `json:"address"`
	State   string       `json:"state"`
	Num1    int          `json:"num1"`
	Num2    int          `json:"num2"`
	String1 string       `json:"string1"`
	String2 string       `json:"string2"`
	Bool    bool         `json:"bool"`
}

func (m *MessageMsg) Instance() MessageMsg {
	return MessageMsg{}
}
