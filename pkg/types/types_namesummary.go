package types

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

var NameParts = []struct {
	Value  coreTypes.Parts
	TSName string
}{
	{coreTypes.Regular, "REGULAR"},
	{coreTypes.Custom, "CUSTOM"},
	{coreTypes.Prefund, "PREFUND"},
	{coreTypes.Baddress, "BADDRESS"},
}

type NameSummary struct {
	NNames     int64                           `json:"nNames"`
	NContracts int64                           `json:"nContracts"`
	NErc20s    int64                           `json:"nErc20s"`
	NErc721s   int64                           `json:"nErc721s"`
	NCustom    int64                           `json:"nCustom"`
	NRegular   int64                           `json:"nRegular"`
	NPrefund   int64                           `json:"nPrefund"`
	NBaddress  int64                           `json:"nBaddress"`
	NDeleted   int64                           `json:"nDeleted"`
	NamesMap   map[base.Address]coreTypes.Name `json:"namesMap"`
	Names      []coreTypes.Name                `json:"names"`
}

func (s *NameSummary) Summarize() {
	s.NNames = int64(len(s.Names))
	for _, name := range s.Names {
		if name.Parts&coreTypes.Regular > 0 {
			s.NRegular++
		}
		if name.Parts&coreTypes.Custom > 0 {
			s.NCustom++
		}
		if name.Parts&coreTypes.Prefund > 0 {
			s.NPrefund++
		}
		if name.Parts&coreTypes.Baddress > 0 {
			s.NBaddress++
		}
		if name.Deleted {
			s.NDeleted++
		}
		if name.IsErc20 {
			s.NErc20s++
		}
		if name.IsErc721 {
			s.NErc721s++
		}
		if name.IsContract {
			s.NContracts++
		}
	}
}

func (s *NameSummary) ShallowCopy() NameSummary {
	return NameSummary{
		NNames:     s.NNames,
		NContracts: s.NContracts,
		NErc20s:    s.NErc20s,
		NErc721s:   s.NErc721s,
		NCustom:    s.NCustom,
		NRegular:   s.NRegular,
		NPrefund:   s.NPrefund,
		NBaddress:  s.NBaddress,
		NDeleted:   s.NDeleted,
	}
}
