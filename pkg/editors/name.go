package editors

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
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

func CoreToName(nameIn coreTypes.Name) Name {
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

func NameToCore(nameIn Name) coreTypes.Name {
	return coreTypes.Name{
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
