package types

import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

type IndexContainer struct {
	coreTypes.ChunkStats
	Items      []coreTypes.ChunkStats `json:"items"`
	NItems     int                    `json:"nItems"`
	Sorts      sdk.SortSpec           `json:"sort"`
	LastUpdate time.Time              `json:"lastUpdate"`
	Chain      string                 `json:"chain"`
}

func NewIndexContainer(chain string, items []coreTypes.ChunkStats) IndexContainer {
	latest := utils.MustGetLatestFileTime(config.PathToIndex(chain))
	return IndexContainer{
		Items: items,
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
		LastUpdate: latest,
		Chain:      chain,
	}
}

func (s *IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *IndexContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(config.PathToIndex(s.Chain))
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *IndexContainer) ShallowCopy() Containerer {
	return &IndexContainer{
		NItems:     s.NItems,
		ChunkStats: s.ChunkStats,
		LastUpdate: s.LastUpdate,
		Chain:      s.Chain,
	}
}

func (s *IndexContainer) Summarize() {
	s.NItems = len(s.Items)
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
}
