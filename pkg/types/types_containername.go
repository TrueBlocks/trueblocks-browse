// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"sort"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type NameContainer struct {
	NContracts uint64           `json:"nContracts"`
	NCustom    uint64           `json:"nCustom"`
	NDeleted   uint64           `json:"nDeleted"`
	NErc20s    uint64           `json:"nErc20s"`
	NErc721s   uint64           `json:"nErc721s"`
	NPrefund   uint64           `json:"nPrefund"`
	NRegular   uint64           `json:"nRegular"`
	NSystem    uint64           `json:"nSystem"`
	SizeOnDisc uint64           `json:"sizeOnDisc"`
	Items      []coreTypes.Name `json:"items"`
	NItems     uint64           `json:"nItems"`
	Chain      string           `json:"chain"`
	LastUpdate time.Time        `json:"lastUpdate"`
	// EXISTING_CODE
	NamesMap map[base.Address]coreTypes.Name `json:"namesMap"`
	// EXISTING_CODE
}

func NewNameContainer(chain string, itemsIn []coreTypes.Name) NameContainer {
	ret := NameContainer{
		Items: make([]coreTypes.Name, 0, len(itemsIn)),
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getNameReload()
	// EXISTING_CODE
	ret.Items = itemsIn
	ret.NamesMap = make(map[base.Address]coreTypes.Name)
	for _, name := range ret.Items {
		ret.NamesMap[name.Address] = name
	}
	sort.Slice(ret.Items, func(i, j int) bool {
		return compare(ret.Items[i], ret.Items[j])
	})
	// EXISTING_CODE
	return ret
}

func (s *NameContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *NameContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getNameReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *NameContainer) ShallowCopy() Containerer {
	return &NameContainer{
		NContracts: s.NContracts,
		NCustom:    s.NCustom,
		NDeleted:   s.NDeleted,
		NErc20s:    s.NErc20s,
		NErc721s:   s.NErc721s,
		NPrefund:   s.NPrefund,
		NRegular:   s.NRegular,
		NSystem:    s.NSystem,
		SizeOnDisc: s.SizeOnDisc,
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *NameContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	for _, name := range s.Items {
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
	chain := "mainnet"
	customPath := filepath.Join(coreConfig.MustGetPathToChainConfig(chain), string(names.DatabaseCustom))
	s.SizeOnDisc = uint64(file.FileSize(customPath))
	regularPath := filepath.Join(coreConfig.MustGetPathToChainConfig(chain), string(names.DatabaseRegular))
	s.SizeOnDisc += uint64(file.FileSize(regularPath))
	// EXISTING_CODE
}

func (s *NameContainer) getNameReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = file.MustGetLatestFileTime(coreConfig.MustGetPathToChainConfig(s.Chain))
	reload = ret != s.LastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
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

// EXISTING_CODE
