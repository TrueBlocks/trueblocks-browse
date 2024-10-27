package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type ManifestContainer struct {
	coreTypes.Manifest `json:",inline"`
	NBlooms            int       `json:"nBlooms"`
	BloomsSize         int       `json:"bloomsSize"`
	NIndexes           int       `json:"nIndexes"`
	IndexSize          int       `json:"indexSize"`
	NItems             int       `json:"nItems"`
	LastUpdate         time.Time `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewManifestContainer(chain string, manifest coreTypes.Manifest) ManifestContainer {
	latest := utils.MustGetLatestFileTime(config.PathToManifest(chain))
	ret := ManifestContainer{
		Manifest:   manifest,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Chain = chain
	// EXISTING_CODE
	return ret
}

func (s *ManifestContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ManifestContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(config.PathToManifest(s.Chain))
	if force || latest != s.LastUpdate {
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
		NBlooms:    s.NBlooms,
		BloomsSize: s.BloomsSize,
		NIndexes:   s.NIndexes,
		IndexSize:  s.IndexSize,
		NItems:     s.NItems,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *ManifestContainer) Summarize() {
	// EXISTING_CODE
	s.NItems = len(s.Chunks)
	for _, item := range s.Chunks {
		s.NBlooms++
		s.BloomsSize += int(item.BloomSize)
		s.NIndexes++
		s.IndexSize += int(item.IndexSize)
	}
	// EXISTING_CODE
}

func X() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
