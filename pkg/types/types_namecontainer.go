package types

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type NameContainer struct {
	Names      []coreTypes.Name                `json:"names"`
	SizeOnDisc int                             `json:"sizeOnDisc"`
	NamesMap   map[base.Address]coreTypes.Name `json:"namesMap"`
	NItems     int                             `json:"nItems"`
	NContracts int                             `json:"nContracts"`
	NErc20s    int                             `json:"nErc20s"`
	NErc721s   int                             `json:"nErc721s"`
	NCustom    int                             `json:"nCustom"`
	NRegular   int                             `json:"nRegular"`
	NPrefund   int                             `json:"nPrefund"`
	NBaddress  int                             `json:"nBaddress"`
	NDeleted   int                             `json:"nDeleted"`
}

func (s *NameContainer) Summarize() {
	s.NItems = len(s.Names)
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

func (s *NameContainer) ShallowCopy() NameContainer {
	return NameContainer{
		NItems:     s.NItems,
		SizeOnDisc: s.SizeOnDisc,
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
