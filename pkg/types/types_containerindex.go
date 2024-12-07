// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type IndexContainer struct {
	Chain      string `json:"chain"`
	ChunkStats `json:",inline"`
	Items      []ChunkStats `json:"items"`
	NItems     uint64       `json:"nItems"`
	Updater    sdk.Updater  `json:"updater"`
	Sorts      sdk.SortSpec `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewIndexContainer(chain string, chunkstats []ChunkStats) IndexContainer {
	// EXISTING_CODE
	itemsIn := chunkstats
	// EXISTING_CODE
	ret := IndexContainer{
		Items:      itemsIn,
		NItems:     uint64(len(itemsIn)),
		ChunkStats: chunkstats[0].ShallowCopy(),
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
		Updater: NewIndexUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
	// EXISTING_CODE
	return ret
}

func NewIndexUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Path: walk.GetRootPathFromCacheType(chain, walk.Index_Bloom), Type: sdk.Folder},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("indexes", items)
	if reset {
		u.Reset()
	}
	return u
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
		NItems:     s.NItems,
		Updater:    s.Updater,
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

func (s *IndexContainer) CollateAndFilter(filter *Filter) interface{} {
	s.Clear()

	logger.InfoBM("CollateAndFilter:", filter.String())
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
		if s.passesFilter(item, filter) {
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

func (s *IndexContainer) Sort() (err error) {
	// EXISTING_CODE
	err = sdk.SortChunkStats(s.Items, s.Sorts)
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
