package editors

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type Name struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Tags     string `json:"tags"`
	Source   string `json:"source"`
	Symbol   string `json:"symbol"`
	Decimals uint64 `json:"decimals"`
	Deleted  bool   `json:"deleted,omitempty"`
}

func CoreToName(nameIn types.Name) Name {
	return Name{
		Address:  nameIn.Address.Hex(),
		Name:     nameIn.Name,
		Tags:     nameIn.Tags,
		Source:   nameIn.Source,
		Symbol:   nameIn.Symbol,
		Decimals: nameIn.Decimals,
		Deleted:  nameIn.Deleted,
	}
}

func NameToCore(nameIn Name) types.Name {
	return types.Name{
		Address:  base.HexToAddress(nameIn.Address),
		Name:     nameIn.Name,
		Tags:     nameIn.Tags,
		Source:   nameIn.Source,
		Symbol:   nameIn.Symbol,
		Decimals: nameIn.Decimals,
		Deleted:  nameIn.Deleted,
	}
}

func (s *Name) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}
