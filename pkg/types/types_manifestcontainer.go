package types

import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type ManifestContainer struct {
	coreTypes.Manifest `json:",inline"`
	NItems             int       `json:"nItems"`
	NBlooms            int       `json:"nBlooms"`
	BloomsSize         int       `json:"bloomsSize"`
	NIndexes           int       `json:"nIndexes"`
	IndexSize          int       `json:"indexSize"`
	LastUpdate         time.Time `json:"lastUpdate"`
}

func NewManifestContainer(chain string, manifest coreTypes.Manifest) ManifestContainer {
	latest := utils.MustGetLatestFileTime(config.PathToManifest(chain))
	ret := ManifestContainer{
		Manifest:   manifest,
		LastUpdate: latest,
	}
	ret.Chain = chain
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
		NItems:     s.NItems,
		NBlooms:    s.NBlooms,
		BloomsSize: s.BloomsSize,
		NIndexes:   s.NIndexes,
		IndexSize:  s.IndexSize,
		LastUpdate: s.LastUpdate,
	}
}

func (s *ManifestContainer) Summarize() {
	s.NItems = len(s.Chunks)
	for _, item := range s.Chunks {
		s.NBlooms++
		s.BloomsSize += int(item.BloomSize)
		s.NIndexes++
		s.IndexSize += int(item.IndexSize)
	}
}
