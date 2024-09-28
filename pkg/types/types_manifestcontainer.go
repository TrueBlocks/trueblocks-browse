package types

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
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
	LastUpdate         time.Time               `json:"lastUpdate"`
}

func NewManifestContainer(chain string, manifest coreTypes.Manifest) ManifestContainer {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToIndex(chain), "finalized"))
	ret := ManifestContainer{
		Manifest:   manifest,
		Items:      manifest.Chunks,
		LastUpdate: latest,
	}
	ret.Chain = chain
	ret.Summarize()

	return ret
}

func (s *ManifestContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ManifestContainer) NeedsUpdate() bool {
	latest := utils.MustGetLatestFileTime(config.PathToManifest(s.Chain))
	if latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *ManifestContainer) ShallowCopy() Containerer {
	return &ManifestContainer{
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
		LastUpdate:   s.LastUpdate,
	}
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
