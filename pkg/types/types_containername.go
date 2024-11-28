// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"sort"
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var namesChain = "mainnet"

// EXISTING_CODE

type NameContainer struct {
	Chain      string          `json:"chain"`
	Items      []Name          `json:"items"`
	NContracts uint64          `json:"nContracts"`
	NCustom    uint64          `json:"nCustom"`
	NDeleted   uint64          `json:"nDeleted"`
	NErc20s    uint64          `json:"nErc20s"`
	NErc721s   uint64          `json:"nErc721s"`
	NItems     uint64          `json:"nItems"`
	NPrefund   uint64          `json:"nPrefund"`
	NRegular   uint64          `json:"nRegular"`
	NSystem    uint64          `json:"nSystem"`
	SizeOnDisc uint64          `json:"sizeOnDisc"`
	Updater    updater.Updater `json:"updater"`
	Sorts      sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewNameContainer(chain string, itemsIn []Name) NameContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := NameContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewNameUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = namesChain // all names are on mainnet
	// EXISTING_CODE
	return ret
}

func NewNameUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: coreConfig.MustGetPathToChainConfig(namesChain), Type: updater.Folder},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("names", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *NameContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *NameContainer) GetItems() interface{} {
	return s.Items
}

func (s *NameContainer) SetItems(items interface{}) {
	s.Items = items.([]Name)
}

func (s *NameContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *NameContainer) ShallowCopy() Containerer {
	ret := &NameContainer{
		Chain:      s.Chain,
		NContracts: s.NContracts,
		NCustom:    s.NCustom,
		NDeleted:   s.NDeleted,
		NErc20s:    s.NErc20s,
		NErc721s:   s.NErc721s,
		NItems:     s.NItems,
		NPrefund:   s.NPrefund,
		NRegular:   s.NRegular,
		NSystem:    s.NSystem,
		SizeOnDisc: s.SizeOnDisc,
		Updater:    s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *NameContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	s.NContracts = 0
	s.NCustom = 0
	s.NDeleted = 0
	s.NErc20s = 0
	s.NErc721s = 0
	s.NPrefund = 0
	s.NRegular = 0
	s.NSystem = 0
	// EXISTING_CODE
}

func (s *NameContainer) passesFilter(item *Name, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		ss := strings.ToLower(filter.Criteria)
		n := strings.ToLower(item.Name)
		a := strings.ToLower(item.Address.Hex())
		t := strings.ToLower(item.Tags)
		c1 := strings.Contains(n, ss)
		c2 := strings.Contains(a, ss)
		c3 := strings.Contains(t, ss)
		ret = c1 || c2 || c3
		// EXISTING_CODE
	}
	return
}

func (s *NameContainer) Accumulate(item *Name) {
	s.NItems++
	// EXISTING_CODE
	if item.Parts&coreTypes.Regular > 0 {
		s.NRegular++
	}
	if item.Parts&coreTypes.Custom > 0 {
		s.NCustom++
	}
	if item.Parts&coreTypes.Prefund > 0 {
		s.NPrefund++
	}
	if item.Parts&coreTypes.Baddress > 0 {
		s.NSystem++
	}
	if item.Deleted {
		s.NDeleted++
	}
	if item.IsErc20 {
		s.NErc20s++
	}
	if item.IsErc721 {
		s.NErc721s++
	}
	if item.IsContract {
		s.NContracts++
	}
	// EXISTING_CODE
}

func (s *NameContainer) Finalize() {
	// EXISTING_CODE
	chain := namesChain
	customPath := filepath.Join(coreConfig.MustGetPathToChainConfig(chain), string(names.DatabaseCustom))
	regularPath := filepath.Join(coreConfig.MustGetPathToChainConfig(chain), string(names.DatabaseRegular))
	s.SizeOnDisc = uint64(file.FileSize(customPath)) + uint64(file.FileSize(regularPath))
	// EXISTING_CODE
}

func (s *NameContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("names") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *Name, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []Name{}
	s.ForEveryItem(func(item *Name, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

func (s *NameContainer) ForEveryItem(process EveryNameFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *NameContainer) Sort() error {
	sort.Slice(s.Items, func(i, j int) bool {
		return compare(s.Items[i], s.Items[j])
	})
	// return sdk.SortNames(s.Items, s.Sorts)
	return nil
}

// EXISTING_CODE
func compare(nameI, nameJ Name) bool {
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
