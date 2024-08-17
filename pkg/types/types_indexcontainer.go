package types

import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type IndexContainer struct {
	coreTypes.ChunkStats
	Items  []coreTypes.ChunkStats `json:"items"`
	NItems uint64                 `json:"nItems"`
}

func (s IndexContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *IndexContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
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
