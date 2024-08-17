package types

import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type ManifestContainer struct {
	coreTypes.Manifest `json:",inline"`
	Items              []coreTypes.ChunkRecord `json:"items"`
	NItems             uint64                  `json:"nItems"`
	LatestUpdate       string                  `json:"latestUpdate"`
	NBlooms            uint64                  `json:"nBlooms"`
	BloomsSize         int64                   `json:"bloomsSize"`
	NIndexes           uint64                  `json:"nIndexes"`
	IndexSize          int64                   `json:"indexSize"`
}

func (s ManifestContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func NewManifestContainer(manifest coreTypes.Manifest) ManifestContainer {
	ret := ManifestContainer{
		Manifest: manifest,
		Items:    manifest.Chunks,
	}
	ret.Summarize()

	return ret
}

func (s *ManifestContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	for _, item := range s.Items {
		s.NBlooms++
		s.BloomsSize += item.BloomSize
		s.NIndexes++
		s.IndexSize += item.IndexSize
	}

}

func (s *ManifestContainer) ShallowCopy() ManifestContainer {
	return ManifestContainer{
		Manifest: coreTypes.Manifest{
			Chain:         s.Manifest.Chain,
			Specification: s.Manifest.Specification,
			Version:       s.Manifest.Version,
		},
		LatestUpdate: s.LatestUpdate,
		NItems:       s.NItems,
		NBlooms:      s.NBlooms,
		BloomsSize:   s.BloomsSize,
		NIndexes:     s.NIndexes,
		IndexSize:    s.IndexSize,
	}
}
