// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type IndexContainer struct {
	Items      []coreTypes.ChunkStats `json:"items"`
	NItems     uint64                 `json:"nItems"`
	Sorts      sdk.SortSpec           `json:"sorts"`
	Chain      string                 `json:"chain"`
	LastUpdate time.Time              `json:"lastUpdate"`
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
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getIndexReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *IndexContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getIndexReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *IndexContainer) ShallowCopy() Containerer {
	return &IndexContainer{
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		ChunkStats: s.ChunkStats,
		Sorts:      s.Sorts,
		// EXISTING_CODE
	}
}

func (s *IndexContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	for _, chunk := range s.Items {
		s.BloomSz += chunk.BloomSz
		s.ChunkSz += chunk.ChunkSz
		s.NAddrs += chunk.NAddrs
		s.NApps += chunk.NApps
		s.NBlocks += chunk.NBlocks
		s.NBlooms += chunk.NBlooms
	}
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

func (s *IndexContainer) getIndexReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = file.MustGetLatestFileTime(coreConfig.PathToIndex(s.Chain))
	reload = ret != s.LastUpdate
	// EXISTING_CODE
	return
}

type EveryChunkStatsFn func(item coreTypes.ChunkStats, data any) bool

func (s *IndexContainer) ForEveryChunkStats(process EveryChunkStatsFn, data any) bool {
	for _, item := range s.Items {
		if !process(item, data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE
