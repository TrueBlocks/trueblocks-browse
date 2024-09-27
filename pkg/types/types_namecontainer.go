package types

import (
	"encoding/json"
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
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

func (a *NameContainer) String() string {
	bytes, _ := json.Marshal(a)
	return string(bytes)
}

func (s *NameContainer) NeedsUpdate() bool {
	e := s.ShallowCopy()
	s.Summarize()
	return (e.NDeleted != s.NDeleted ||
		e.NCustom != s.NCustom ||
		e.SizeOnDisc != s.SizeOnDisc ||
		e.NItems != s.NItems ||
		e.NRegular != s.NRegular ||
		e.NContracts != s.NContracts ||
		e.NErc20s != s.NErc20s ||
		e.NErc721s != s.NErc721s ||
		e.NPrefund != s.NPrefund ||
		e.NBaddress != s.NBaddress)
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

func (s *NameContainer) Summarize() {
	chain := "mainnet"
	customPath := filepath.Join(config.MustGetPathToChainConfig(chain), string(names.DatabaseCustom))
	s.SizeOnDisc = int(file.FileSize(customPath))
	regularPath := filepath.Join(config.MustGetPathToChainConfig(chain), string(names.DatabaseRegular))
	s.SizeOnDisc += int(file.FileSize(regularPath))

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
