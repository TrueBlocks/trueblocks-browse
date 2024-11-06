// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type StatusContainer struct {
	NBytes           uint64 `json:"nBytes"`
	NFiles           uint64 `json:"nFiles"`
	NFolders         uint64 `json:"nFolders"`
	coreTypes.Status `json:",inline"`
	LastUpdate       int64 `json:"lastUpdate"`
	// EXISTING_CODE
	Items  []coreTypes.CacheItem `json:"items"`
	NItems uint64                `json:"nItems"`
	// EXISTING_CODE
}

func NewStatusContainer(chain string, status *coreTypes.Status) StatusContainer {
	ret := StatusContainer{
		Status: *status,
	}
	ret.Chain = chain
	ret.LastUpdate, _ = ret.getStatusReload()
	// EXISTING_CODE
	ret.LastUpdate = time.Now().Unix()
	ret.Items = status.Caches
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
		DebugInts("reload Status", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *StatusContainer) ShallowCopy() Containerer {
	ret := &StatusContainer{
		NBytes:     s.NBytes,
		NFiles:     s.NFiles,
		NFolders:   s.NFolders,
		Status:     s.Status.ShallowCopy(),
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		NItems: s.NItems,
		// EXISTING_CODE
	}
	ret.Chain = s.Chain
	return ret
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

func (s *StatusContainer) getStatusReload() (ret int64, reload bool) {
	// EXISTING_CODE
	ret = time.Now().Unix()
	reload = ret > s.LastUpdate+(60*2) // every two minutes
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
