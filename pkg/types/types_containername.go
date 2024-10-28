package types

// EXISTING_CODE
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

type NameItemType = coreTypes.Name
type NameInputType = map[base.Address]coreTypes.Name

// EXISTING_CODE

type NameContainer struct {
	NContracts uint64         `json:"nContracts"`
	NCustom    uint64         `json:"nCustom"`
	NDeleted   uint64         `json:"nDeleted"`
	NErc20s    uint64         `json:"nErc20s"`
	NErc721s   uint64         `json:"nErc721s"`
	NPrefund   uint64         `json:"nPrefund"`
	NRegular   uint64         `json:"nRegular"`
	NSystem    uint64         `json:"nSystem"`
	SizeOnDisc uint64         `json:"sizeOnDisc"`
	Items      []NameItemType `json:"items"`
	NItems     uint64         `json:"nItems"`
	Chain      string         `json:"chain"`
	LastUpdate time.Time      `json:"lastUpdate"`
	// EXISTING_CODE
	NamesMap map[base.Address]coreTypes.Name `json:"namesMap"`
	// EXISTING_CODE
}

func NewNameContainer(chain string, itemsIn NameInputType) NameContainer {
	latest, _ := getNameReload(chain, time.Time{})
	ret := NameContainer{
		Items:      make([]NameItemType, 0, len(itemsIn)),
		Chain:      chain,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.NamesMap = itemsIn
	for _, name := range ret.NamesMap {
		ret.Items = append(ret.Items, name)
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
	latest, reload := getNameReload(s.Chain, s.LastUpdate)
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
	customPath := filepath.Join(config.MustGetPathToChainConfig(chain), string(names.DatabaseCustom))
	s.SizeOnDisc = uint64(file.FileSize(customPath))
	regularPath := filepath.Join(config.MustGetPathToChainConfig(chain), string(names.DatabaseRegular))
	s.SizeOnDisc += uint64(file.FileSize(regularPath))
	// EXISTING_CODE
}

func getNameReload(chain string, lastUpdate time.Time) (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = utils.MustGetLatestFileTime(config.MustGetPathToChainConfig(chain))
	reload = ret != lastUpdate
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
