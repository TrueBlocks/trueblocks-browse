package types

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually this will get put back into Core.

type ManifestSummary struct {
	coreTypes.Manifest `json:",inline"`
	LatestUpdate       string `json:"latestUpdate"`
	NBlooms            uint64 `json:"nBlooms"`
	BloomsSize         int64  `json:"bloomsSize"`
	NIndexes           uint64 `json:"nIndexes"`
	IndexSize          int64  `json:"indexSize"`
}

func NewManifestEx(manifest coreTypes.Manifest) ManifestSummary {
	ret := ManifestSummary{
		Manifest: manifest,
	}

	for _, chunk := range manifest.Chunks {
		ret.NBlooms++
		ret.BloomsSize += chunk.BloomSize
		ret.NIndexes++
		ret.IndexSize += chunk.IndexSize
	}

	return ret
}

func (s *ManifestSummary) ShallowCopy() ManifestSummary {
	return ManifestSummary{
		Manifest: coreTypes.Manifest{
			Chain:         s.Manifest.Chain,
			Specification: s.Manifest.Specification,
			Version:       s.Manifest.Version,
		},
		LatestUpdate: s.LatestUpdate,
		NBlooms:      s.NBlooms,
		BloomsSize:   s.BloomsSize,
		NIndexes:     s.NIndexes,
		IndexSize:    s.IndexSize,
	}
}
