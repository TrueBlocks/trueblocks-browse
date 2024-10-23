package types

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type NamesContainer struct {
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
	NSystem    int                             `json:"nSystem"`
	NDeleted   int                             `json:"nDeleted"`
	LastUpdate time.Time                       `json:"lastUpdate"`
	Chain      string                          `json:"chain"`
}

func NewNamesContainer(chain string, namesMap map[base.Address]coreTypes.Name) NamesContainer {
	latest := utils.MustGetLatestFileTime(config.MustGetPathToChainConfig(chain))
	return NamesContainer{
		Names:      make([]coreTypes.Name, 0),
		NamesMap:   namesMap,
		LastUpdate: latest,
		Chain:      chain,
	}
}

func (a *NamesContainer) String() string {
	bytes, _ := json.Marshal(a)
	return string(bytes)
}

func (s *NamesContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(config.MustGetPathToChainConfig(s.Chain))
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *NamesContainer) ShallowCopy() Containerer {
	return &NamesContainer{
		NItems:     s.NItems,
		SizeOnDisc: s.SizeOnDisc,
		NContracts: s.NContracts,
		NErc20s:    s.NErc20s,
		NErc721s:   s.NErc721s,
		NCustom:    s.NCustom,
		NRegular:   s.NRegular,
		NPrefund:   s.NPrefund,
		NSystem:    s.NSystem,
		NDeleted:   s.NDeleted,
		LastUpdate: s.LastUpdate,
		Chain:      s.Chain,
	}
}

func (s *NamesContainer) Summarize() {
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
