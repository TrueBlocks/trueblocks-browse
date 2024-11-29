// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type ProjectContainer struct {
	Chain       string             `json:"chain"`
	HistorySize uint64             `json:"historySize"`
	Items       []HistoryContainer `json:"items"`
	NAbis       uint64             `json:"nAbis"`
	NCaches     uint64             `json:"nCaches"`
	NIndexes    uint64             `json:"nIndexes"`
	NItems      uint64             `json:"nItems"`
	NManifests  uint64             `json:"nManifests"`
	NMonitors   uint64             `json:"nMonitors"`
	NNames      uint64             `json:"nNames"`
	Updater     updater.Updater    `json:"updater"`
	Sorts       sdk.SortSpec       `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewProjectContainer(chain string, itemsIn []HistoryContainer) ProjectContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := ProjectContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{"address"},
			Order:  []sdk.SortOrder{sdk.Asc},
		},
		Updater: NewProjectUpdater(chain, itemsIn),
	}
	// EXISTING_CODE
	ret.Chain = chain
	// EXISTING_CODE
	return ret
}

func NewProjectUpdater(chain string, itemsIn []HistoryContainer, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: coreConfig.MustGetPathToChainConfig(namesChain), Type: updater.Folder},
	}
	for _, item := range itemsIn {
		path := coreMonitor.PathToMonitorFile(chain, item.Address)
		item := updater.UpdaterItem{Path: path, Type: updater.FileSize}
		items = append(items, item)
	}
	// EXISTING_CODE
	u, _ := updater.NewUpdater("project", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *ProjectContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ProjectContainer) GetItems() interface{} {
	return s.Items
}

func (s *ProjectContainer) SetItems(items interface{}) {
	s.Items = items.([]HistoryContainer)
}

func (s *ProjectContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *ProjectContainer) ShallowCopy() Containerer {
	ret := &ProjectContainer{
		Chain:       s.Chain,
		HistorySize: s.HistorySize,
		NAbis:       s.NAbis,
		NCaches:     s.NCaches,
		NIndexes:    s.NIndexes,
		NItems:      s.NItems,
		NManifests:  s.NManifests,
		NMonitors:   s.NMonitors,
		NNames:      s.NNames,
		Updater:     s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *ProjectContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ProjectContainer) passesFilter(item *HistoryContainer, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *ProjectContainer) Accumulate(item *HistoryContainer) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ProjectContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ProjectContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("project") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *HistoryContainer, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []HistoryContainer{}
	s.ForEveryItem(func(item *HistoryContainer, data any) bool {
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

func (s *ProjectContainer) ForEveryItem(process EveryHistoryContainerFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *ProjectContainer) Sort() (err error) {
	// EXISTING_CODE
	sort.Slice(s.Items, func(i, j int) bool {
		return s.Items[i].Address.Hex() < s.Items[j].Address.Hex()
	})
	// TODO: Sorting?
	// return sdk.SortHistoryContainers(s.Items, s.Sorts)
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
