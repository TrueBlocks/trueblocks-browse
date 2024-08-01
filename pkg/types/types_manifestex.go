package types

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually this will get put back into Core.

type ManifestEx struct {
	coreTypes.Manifest `json:",inline"`
	LatestUpdate       string `json:"latestUpdate"`
	NBlooms            uint64 `json:"nBlooms"`
	BloomsSize         int64  `json:"bloomsSize"`
	NIndexes           uint64 `json:"nIndexes"`
	IndexSize          int64  `json:"indexSize"`
}

func NewManifestEx(manifest coreTypes.Manifest) ManifestEx {
	ret := ManifestEx{
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
