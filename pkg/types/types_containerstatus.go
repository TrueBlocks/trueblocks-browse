package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/version"
)

type StatusItemType = coreTypes.CacheItem
type StatusInputType = []coreTypes.Status

// EXISTING_CODE

type StatusContainer struct {
	NBytes     uint64           `json:"nBytes"`
	NFiles     uint64           `json:"nFiles"`
	NFolders   uint64           `json:"nFolders"`
	Items      []StatusItemType `json:"items"`
	NItems     uint64           `json:"nItems"`
	LastUpdate time.Time        `json:"lastUpdate"`
	// EXISTING_CODE
	coreTypes.Status `json:",inline"`
	// EXISTING_CODE
}

func NewStatusContainer(chain string, itemsIn StatusInputType) StatusContainer {
	latest, _ := getStatusReload(chain, time.Time{})
	ret := StatusContainer{
		Items:      make([]StatusItemType, 0, len(itemsIn)),
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Chain = chain
	ret.LastUpdate = time.Now()
	ret.Status = itemsIn[0]
	// TODO: This is a hack. We need to get the version from the core
	ret.Version = version.LibraryVersion
	ret.Items = itemsIn[0].Caches
	// EXISTING_CODE
	return ret
}

func (s *StatusContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *StatusContainer) NeedsUpdate(force bool) bool {
	latest, reload := getStatusReload(s.Chain, s.LastUpdate)
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *StatusContainer) ShallowCopy() Containerer {
	return &StatusContainer{
		NBytes:     s.NBytes,
		NFiles:     s.NFiles,
		NFolders:   s.NFolders,
		NItems:     s.NItems,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		Status: s.Status.ShallowCopy(),
		// EXISTING_CODE
	}
}

func (s *StatusContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	for _, cache := range s.Caches {
		s.NFolders += cache.NFolders
		s.NFiles += cache.NFiles
		s.NBytes += uint64(cache.SizeInBytes)
	}
	// EXISTING_CODE
}

func getStatusReload(chain string, lastUpdate time.Time) (ret time.Time, reload bool) {
	// EXISTING_CODE
	_ = chain
	ret = time.Now()
	reload = ret.After(lastUpdate.Add(time.Minute * 2)) // every two minutes
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
