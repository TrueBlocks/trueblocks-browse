// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

// -------------------------------------------------------------------
type IndexContainer struct {
	Items      []coreTypes.ChunkStats `json:"items"`
	NItems     uint64                 `json:"nItems"`
	Sorts      sdk.SortSpec           `json:"sorts"`
	Chain      string                 `json:"chain"`
	LastUpdate int64                  `json:"lastUpdate"`
	// EXISTING_CODE
	coreTypes.ChunkStats
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func NewIndexContainer(chain string, itemsIn []coreTypes.ChunkStats) IndexContainer {
	ret := IndexContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getIndexReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

// -------------------------------------------------------------------
func (s *IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

// -------------------------------------------------------------------
func (s *IndexContainer) GetItems() interface{} {
	return s.Items
}

// -------------------------------------------------------------------
func (s *IndexContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.ChunkStats)
}

// -------------------------------------------------------------------
func (s *IndexContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getIndexReload()
	if force || reload {
		DebugInts("index", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

// -------------------------------------------------------------------
func (s *IndexContainer) ShallowCopy() Containerer {
	ret := &IndexContainer{
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		ChunkStats: s.ChunkStats,
		Sorts:      s.Sorts,
		// EXISTING_CODE
	}
	return ret
}

// -------------------------------------------------------------------
func (s *IndexContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	s.BloomSz = 0
	s.ChunkSz = 0
	s.NAddrs = 0
	s.NApps = 0
	s.NBlocks = 0
	s.NBlooms = 0
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *IndexContainer) passesFilter(item *coreTypes.ChunkStats, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

// -------------------------------------------------------------------
func (s *IndexContainer) Accumulate(item *coreTypes.ChunkStats) {
	s.NItems++
	// EXISTING_CODE
	s.BloomSz += item.BloomSz
	s.ChunkSz += item.ChunkSz
	s.NAddrs += item.NAddrs
	s.NApps += item.NApps
	s.NBlocks += item.NBlocks
	s.NBlooms += item.NBlooms
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *IndexContainer) Finalize() {
	// EXISTING_CODE
	if s.NBlocks > 0 {
		s.AddrsPerBlock = float64(s.NAddrs) / float64(s.NBlocks)
	}
	if s.NAddrs > 0 {
		s.AppsPerAddr = float64(s.NApps) / float64(s.NAddrs)
	}
	if s.NBlocks > 0 {
		s.AppsPerBlock = float64(s.NApps) / float64(s.NBlocks)
	}
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *IndexContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("indexes") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *coreTypes.ChunkStats, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []coreTypes.ChunkStats{}
	s.ForEveryItem(func(item *coreTypes.ChunkStats, data any) bool {
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

// -------------------------------------------------------------------
func (s *IndexContainer) getIndexReload() (ret int64, reload bool) {
	// EXISTING_CODE
	tm := file.MustGetLatestFileTime(coreConfig.PathToIndex(s.Chain))
	ret = tm.Unix()
	reload = ret > s.LastUpdate
	// EXISTING_CODE
	return
}

// -------------------------------------------------------------------
type EveryChunkStatsFn func(item *coreTypes.ChunkStats, data any) bool

// -------------------------------------------------------------------
func (s *IndexContainer) ForEveryItem(process EveryChunkStatsFn, data any) bool {
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
// class:         Index
// lower:         index
// routeLabel:    Indexes
// routeLower:    indexes
// embedName:
// embedType:     .
// otherName:
// otherType:     .
// itemName:      ChunkStats
// itemType:      coreTypes.ChunkStats
// inputType:     coreTypes.ChunkStats
// hasEmbed:      false
// hasItems:      true
// hasOther:      false
// hasSorts:      true
// initChain:     false
// isEditable:    false
// needsChain:    true
// needsLoad:     true
// needsSdk:      true
