package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type ModifyData struct {
	Operation string       `json:"operation"`
	Filename  string       `json:"filename"`
	Address   base.Address `json:"address"`
	Value     string       `json:"value"`
}
