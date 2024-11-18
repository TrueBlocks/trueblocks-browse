// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
)

// EXISTING_CODE

// -------------------------------------------------------------------
type ProjectContainer struct {
	HistorySize uint64             `json:"historySize"`
	NAbis       uint64             `json:"nAbis"`
	NCaches     uint64             `json:"nCaches"`
	NIndexes    uint64             `json:"nIndexes"`
	NManifests  uint64             `json:"nManifests"`
	NMonitors   uint64             `json:"nMonitors"`
	NNames      uint64             `json:"nNames"`
	Items       []HistoryContainer `json:"items"`
	NItems      uint64             `json:"nItems"`
	Chain       string             `json:"chain"`
	LastUpdate  int64              `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func NewProjectContainer(chain string, itemsIn []HistoryContainer) ProjectContainer {
	ret := ProjectContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Chain:  chain,
	}
	ret.LastUpdate, _ = ret.getProjectReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

// -------------------------------------------------------------------
func (s *ProjectContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

// -------------------------------------------------------------------
func (s *ProjectContainer) GetItems() interface{} {
	return s.Items
}

// -------------------------------------------------------------------
func (s *ProjectContainer) SetItems(items interface{}) {
	s.Items = items.([]HistoryContainer)
}

// -------------------------------------------------------------------
func (s *ProjectContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getProjectReload()
	if force || reload {
		DebugInts("project", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

// -------------------------------------------------------------------
func (s *ProjectContainer) ShallowCopy() Containerer {
	ret := &ProjectContainer{
		HistorySize: s.HistorySize,
		NAbis:       s.NAbis,
		NCaches:     s.NCaches,
		NIndexes:    s.NIndexes,
		NManifests:  s.NManifests,
		NMonitors:   s.NMonitors,
		NNames:      s.NNames,
		NItems:      s.NItems,
		Chain:       s.Chain,
		LastUpdate:  s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

// -------------------------------------------------------------------
func (s *ProjectContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *ProjectContainer) passesFilter(item *HistoryContainer, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

// -------------------------------------------------------------------
func (s *ProjectContainer) Accumulate(item *HistoryContainer) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *ProjectContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
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
	// do nothing - summaries are already calculated
	// EXISTING_CODE

	return filtered
}

// -------------------------------------------------------------------
func (s *ProjectContainer) getProjectReload() (ret int64, reload bool) {
	// EXISTING_CODE
	var totalSize int64 = 0
	_ = s.ForEveryItem(func(item *HistoryContainer, data any) bool {
		fn := coreMonitor.PathToMonitorFile(s.Chain, item.Address)
		fs := file.FileSize(fn)
		totalSize += fs
		return true
	}, nil)
	if totalSize > s.LastUpdate {
		reload = true
		ret = totalSize
	}
	// EXISTING_CODE
	return
}

// -------------------------------------------------------------------
type EveryHistoryContainerFn func(item *HistoryContainer, data any) bool

// -------------------------------------------------------------------
func (s *ProjectContainer) ForEveryItem(process EveryHistoryContainerFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Project
// lower:         project
// routeLabel:    Project
// routeLower:    project
// embedName:
// embedType:     .
// otherName:
// otherType:     .
// itemName:      HistoryContainer
// itemType:      HistoryContainer
// inputType:     HistoryContainer
// hasItems:      true
// hasEmbed:      false
// hasOther:      false
// hasSorts:      false
// initChain:     false
// isEditable:    false
// needsChain:    true
// needsLoad:     true
// needsSdk:      false
