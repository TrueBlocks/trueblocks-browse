package types

import (
	"encoding/json"
	"path/filepath"
	"sort"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type NameContainer struct {
	Names      []coreTypes.Name                `json:"names"`
	NamesMap   map[base.Address]coreTypes.Name `json:"namesMap"`
	NItems     int                             `json:"nItems"`
	NDeleted   int                             `json:"nDeleted"`
	NContracts int                             `json:"nContracts"`
	NErc20s    int                             `json:"nErc20s"`
	NErc721s   int                             `json:"nErc721s"`
	NCustom    int                             `json:"nCustom"`
	NRegular   int                             `json:"nRegular"`
	NPrefund   int                             `json:"nPrefund"`
	NSystem    int                             `json:"nSystem"`
	LastUpdate time.Time                       `json:"lastUpdate"`
	SizeOnDisc int                             `json:"sizeOnDisc"`
	Chain      string                          `json:"chain"`
}

func NewNamesContainer(chain string, namesMap map[base.Address]coreTypes.Name) NameContainer {
	latest := utils.MustGetLatestFileTime(config.MustGetPathToChainConfig(chain))
	ret := NameContainer{
		Names:      make([]coreTypes.Name, 0),
		NamesMap:   namesMap,
		LastUpdate: latest,
		Chain:      chain,
	}
	for _, name := range ret.NamesMap {
		ret.Names = append(ret.Names, name)
	}
	sort.Slice(ret.Names, func(i, j int) bool {
		return compare(ret.Names[i], ret.Names[j])
	})
	return ret
}

func (a *NameContainer) String() string {
	bytes, _ := json.Marshal(a)
	return string(bytes)
}

func (s *NameContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(config.MustGetPathToChainConfig(s.Chain))
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *NameContainer) ShallowCopy() Containerer {
	return &NameContainer{
		NItems:     s.NItems,
		NDeleted:   s.NDeleted,
		NContracts: s.NContracts,
		NErc20s:    s.NErc20s,
		NErc721s:   s.NErc721s,
		NCustom:    s.NCustom,
		NRegular:   s.NRegular,
		NPrefund:   s.NPrefund,
		NSystem:    s.NSystem,
		LastUpdate: s.LastUpdate,
		SizeOnDisc: s.SizeOnDisc,
		Chain:      s.Chain,
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
			s.NSystem++
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

func compare(nameI, nameJ coreTypes.Name) bool {
	ti := nameI.Parts
	if ti == coreTypes.Regular {
		ti = 7
	}
	tj := nameJ.Parts
	if tj == coreTypes.Regular {
		tj = 7
	}
	if ti == tj {
		if nameI.Tags == nameJ.Tags {
			return nameI.Address.Hex() < nameJ.Address.Hex()
		}
		return nameI.Tags < nameJ.Tags
	}
	return ti < tj
}
