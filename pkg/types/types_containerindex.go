package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

type IndexItemType = coreTypes.ChunkStats
type IndexInputType = []coreTypes.ChunkStats

// EXISTING_CODE

type IndexContainer struct {
	Items      []IndexItemType `json:"items"`
	NItems     uint64          `json:"nItems"`
	Chain      string          `json:"chain"`
	LastUpdate time.Time       `json:"lastUpdate"`
	// EXISTING_CODE
	coreTypes.ChunkStats
	Sorts sdk.SortSpec `json:"sort"`
	// EXISTING_CODE
}

func NewIndexContainer(chain string, itemsIn IndexInputType) IndexContainer {
	latest := getLatestIndexDate(chain)
	ret := IndexContainer{
		Items:      make([]IndexItemType, 0, len(itemsIn)),
		Chain:      chain,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Items = itemsIn
	ret.Sorts = sdk.SortSpec{
		Fields: []string{"range"},
		Order:  []sdk.SortOrder{sdk.Dec},
	}
	// EXISTING_CODE
	return ret
}

func (s *IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *IndexContainer) NeedsUpdate(force bool) bool {
	latest := getLatestIndexDate(s.Chain)
	if force || latest != s.LastUpdate {
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

func getLatestIndexDate(chain string) (ret time.Time) {
	// EXISTING_CODE
	ret = utils.MustGetLatestFileTime(config.PathToIndex(chain))
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
