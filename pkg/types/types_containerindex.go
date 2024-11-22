// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type IndexContainer struct {
	Chain   string                 `json:"chain"`
	Updater walk.Updater           `json:"updater"`
	Items   []coreTypes.ChunkStats `json:"items"`
	NItems  uint64                 `json:"nItems"`
	Sorts   sdk.SortSpec           `json:"sorts"`
	// EXISTING_CODE
	coreTypes.ChunkStats
	// EXISTING_CODE
}

func NewIndexContainer(chain string, itemsIn []coreTypes.ChunkStats) IndexContainer {
	ret := IndexContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
		Chain:   chain,
		Updater: NewIndexUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewIndexUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{
		walk.GetRootPathFromCacheType(chain, walk.Index_Bloom),
	}
	updater, _ := walk.NewUpdater("index", paths, walk.TypeFolders)
	// EXISTING_CODE
	return updater
}

func (s *IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *IndexContainer) GetItems() interface{} {
	return s.Items
}

func (s *IndexContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.ChunkStats)
}

func (s *IndexContainer) NeedsUpdate() bool {
	if updater, reload := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *IndexContainer) ShallowCopy() Containerer {
	ret := &IndexContainer{
		Chain:   s.Chain,
		Updater: s.Updater,
		NItems:  s.NItems,
		// EXISTING_CODE
		ChunkStats: s.ChunkStats,
		Sorts:      s.Sorts,
		// EXISTING_CODE
	}
	return ret
}

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

func (s *IndexContainer) passesFilter(item *coreTypes.ChunkStats, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

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

type EveryChunkStatsFn func(item *coreTypes.ChunkStats, data any) bool

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
