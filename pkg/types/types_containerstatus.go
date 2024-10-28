package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/version"
)

// EXISTING_CODE

type StatusContainer struct {
	coreTypes.Status `json:",inline"`
	NBytes           uint64    `json:"nBytes"`
	NFiles           uint64    `json:"nFiles"`
	NFolders         uint64    `json:"nFolders"`
	LastUpdate       time.Time `json:"lastUpdate"`
	// EXISTING_CODE
	Items  []coreTypes.CacheItem `json:"items"`
	NItems uint64                `json:"nItems"`
	// EXISTING_CODE
}

func NewStatusContainer(chain string, itemsIn *coreTypes.Status) StatusContainer {
	ret := StatusContainer{
		Status: *itemsIn,
	}
	ret.LastUpdate, _ = ret.getStatusReload()
	// EXISTING_CODE
	ret.Chain = chain
	ret.LastUpdate = time.Now()
	// TODO: This is a hack. We need to get the version from the core
	ret.Version = version.LibraryVersion
	ret.Items = itemsIn.Caches
	ret.NItems = uint64(len(ret.Items))
	// EXISTING_CODE
	return ret
}

func (s *StatusContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *StatusContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getStatusReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *StatusContainer) ShallowCopy() Containerer {
	return &StatusContainer{
		Status:     s.Status.ShallowCopy(),
		NBytes:     s.NBytes,
		NFiles:     s.NFiles,
		NFolders:   s.NFolders,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		NItems: s.NItems,
		// EXISTING_CODE
	}
}

func (s *StatusContainer) Summarize() {
	// EXISTING_CODE
	s.NItems = uint64(len(s.Items))
	for _, cache := range s.Caches {
		s.NFolders += cache.NFolders
		s.NFiles += cache.NFiles
		s.NBytes += uint64(cache.SizeInBytes)
	}
	// EXISTING_CODE
}

func (s *StatusContainer) getStatusReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = time.Now()
	reload = ret.After(s.LastUpdate.Add(time.Minute * 2)) // every two minutes
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
