package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type ManifestItemType = coreTypes.ChunkRecord
type ManifestInputType = []coreTypes.Manifest

// EXISTING_CODE

type ManifestContainer struct {
	BloomsSize    uint64             `json:"bloomsSize"`
	IndexSize     uint64             `json:"indexSize"`
	NBlooms       uint64             `json:"nBlooms"`
	NIndexes      uint64             `json:"nIndexes"`
	Specification string             `json:"specification"`
	Version       string             `json:"version"`
	Items         []ManifestItemType `json:"items"`
	NItems        uint64             `json:"nItems"`
	Chain         string             `json:"chain"`
	LastUpdate    time.Time          `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewManifestContainer(chain string, itemsIn ManifestInputType) ManifestContainer {
	latest := getLatestManifestDate(chain)
	ret := ManifestContainer{
		Items:      make([]ManifestItemType, 0, len(itemsIn)),
		Chain:      chain,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Specification = itemsIn[0].Specification.String()
	ret.Version = itemsIn[0].Version
	ret.Items = itemsIn[0].Chunks
	// EXISTING_CODE
	return ret
}

func (s *ManifestContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ManifestContainer) NeedsUpdate(force bool) bool {
	latest := getLatestManifestDate(s.Chain)
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *ManifestContainer) ShallowCopy() Containerer {
	return &ManifestContainer{
		BloomsSize:    s.BloomsSize,
		IndexSize:     s.IndexSize,
		NBlooms:       s.NBlooms,
		NIndexes:      s.NIndexes,
		Specification: s.Specification,
		Version:       s.Version,
		NItems:        s.NItems,
		Chain:         s.Chain,
		LastUpdate:    s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *ManifestContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	for _, item := range s.Items {
		s.BloomsSize += uint64(item.BloomSize)
		s.IndexSize += uint64(item.IndexSize)
		s.NBlooms++
		s.NIndexes++
	}
	// EXISTING_CODE
}

func getLatestManifestDate(chain string) (ret time.Time) {
	// EXISTING_CODE
	ret = utils.MustGetLatestFileTime(config.PathToManifest(chain))
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
