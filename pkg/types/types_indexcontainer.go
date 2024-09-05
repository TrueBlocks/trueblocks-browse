package types

import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

type IndexContainer struct {
	coreTypes.ChunkStats
	Items  []coreTypes.ChunkStats `json:"items"`
	NItems int                    `json:"nItems"`
	Sorts  sdk.SortSpec           `json:"sort"`
}

func NewIndexContainer(items []coreTypes.ChunkStats) IndexContainer {
	return IndexContainer{
		Items: items,
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
	}
}

func (s IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
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

func (s *IndexContainer) ShallowCopy() IndexContainer {
	return IndexContainer{
		NItems:     s.NItems,
		ChunkStats: s.ChunkStats,
	}
}
