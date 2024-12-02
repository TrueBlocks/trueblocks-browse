// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type IndexContainer struct {
	Chain      string `json:"chain"`
	ChunkStats `json:",inline"`
	Updater    updater.Updater `json:"updater"`
	Items      []ChunkStats    `json:"items"`
	NItems     uint64          `json:"nItems"`
	Sorts      sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewIndexContainer(chain string, itemsIn []ChunkStats) IndexContainer {
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

func NewIndexUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: walk.GetRootPathFromCacheType(chain, walk.Index_Bloom), Type: updater.Folder},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("indexes", items)
	if reset {
		updater.Reset()
	}
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
	s.Items = items.([]ChunkStats)
}

func (s *IndexContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *IndexContainer) ShallowCopy() Containerer {
	ret := &IndexContainer{
		Chain:      s.Chain,
		ChunkStats: s.ChunkStats.ShallowCopy(),
		Updater:    s.Updater,
		NItems:     s.NItems,
		// EXISTING_CODE
		Sorts: s.Sorts,
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

func (s *IndexContainer) passesFilter(item *ChunkStats, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *IndexContainer) Accumulate(item *ChunkStats) {
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
		s.ForEveryItem(func(item *ChunkStats, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []ChunkStats{}
	s.ForEveryItem(func(item *ChunkStats, data any) bool {
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
