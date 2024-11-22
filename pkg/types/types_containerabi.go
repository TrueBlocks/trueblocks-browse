// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"
	"path/filepath"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type AbiContainer struct {
	Chain         string          `json:"chain"`
	LargestFile   string          `json:"largestFile"`
	MostEvents    string          `json:"mostEvents"`
	MostFunctions string          `json:"mostFunctions"`
	Updater       walk.Updater    `json:"updater"`
	Items         []coreTypes.Abi `json:"items"`
	NItems        uint64          `json:"nItems"`
	Sorts         sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	coreTypes.Abi
	// EXISTING_CODE
}

func NewAbiContainer(chain string, itemsIn []coreTypes.Abi) AbiContainer {
	ret := AbiContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{"isEmpty", "isKnown", "address"},
			Order:  []sdk.SortOrder{sdk.Asc, sdk.Asc, sdk.Asc},
		},
		Chain:   chain,
		Updater: NewAbiUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewAbiUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{
		filepath.Join(coreConfig.PathToCache(chain), "abis"),
		coreConfig.MustGetPathToChainConfig(namesChain),
	}
	updater, _ := walk.NewUpdater("abis", paths, walk.TypeFolders)
	// EXISTING_CODE
	return updater
}

func (s *AbiContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *AbiContainer) GetItems() interface{} {
	return s.Items
}

func (s *AbiContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.Abi)
}

func (s *AbiContainer) NeedsUpdate() bool {
	if updater, reload := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *AbiContainer) ShallowCopy() Containerer {
	ret := &AbiContainer{
		Chain:         s.Chain,
		LargestFile:   s.LargestFile,
		MostEvents:    s.MostEvents,
		MostFunctions: s.MostFunctions,
		Updater:       s.Updater,
		NItems:        s.NItems,
		// EXISTING_CODE
		Abi: s.Abi,
		// EXISTING_CODE
	}
	return ret
}

func (s *AbiContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *AbiContainer) passesFilter(item *coreTypes.Abi, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *AbiContainer) Accumulate(item *coreTypes.Abi) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *AbiContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *AbiContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("abis") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *coreTypes.Abi, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []coreTypes.Abi{}
	s.ForEveryItem(func(item *coreTypes.Abi, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	s.FileSize = 0
	s.NEvents = 0
	s.NFunctions = 0
	var lF comparison
	var mF comparison
	var mE comparison
	for _, fil := range s.Items {
		s.NFunctions += fil.NFunctions
		s.NEvents += fil.NEvents
		s.FileSize += fil.FileSize
		lF.MarkMax(fil.Name, int(fil.FileSize))
		mF.MarkMax(fil.Name, int(fil.NFunctions))
		mE.MarkMax(fil.Name, int(fil.NEvents))
	}
	s.LargestFile = fmt.Sprintf("%s (%d bytes)", lF.Name, lF.Value)
	s.MostFunctions = fmt.Sprintf("%s (%d functions)", mF.Name, mF.Value)
	s.MostEvents = fmt.Sprintf("%s (%d events)", mE.Name, mE.Value)
	// EXISTING_CODE

	return filtered
}

type EveryAbiFn func(item *coreTypes.Abi, data any) bool

func (s *AbiContainer) ForEveryItem(process EveryAbiFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
type comparison struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (c *comparison) MarkMax(name string, value int) {
	if c.Value < value {
		c.Name = name
		c.Value = value
	}
}

func (c *comparison) MarkMin(name string, value int) {
	if c.Value > value {
		c.Name = name
		c.Value = value
	}
}

// EXISTING_CODE
