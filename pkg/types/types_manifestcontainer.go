package types

import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type ManifestContainer struct {
	coreTypes.Manifest `json:",inline"`
	Items              []coreTypes.ChunkRecord `json:"items"`
	NItems             int                     `json:"nItems"`
	LatestUpdate       string                  `json:"latestUpdate"`
	NBlooms            int                     `json:"nBlooms"`
	BloomsSize         int                     `json:"bloomsSize"`
	NIndexes           int                     `json:"nIndexes"`
	IndexSize          int                     `json:"indexSize"`
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
	s.NItems = len(s.Items)
	for _, item := range s.Items {
		s.NBlooms++
		s.BloomsSize += int(item.BloomSize)
		s.NIndexes++
		s.IndexSize += int(item.IndexSize)
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
