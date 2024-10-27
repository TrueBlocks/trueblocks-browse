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
	NItems           int       `json:"nItems"`
	NFolders         int       `json:"nFolders"`
	NFiles           int       `json:"nFiles"`
	NBytes           int       `json:"nBytes"`
	LastUpdate       time.Time `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewStatusContainer(chain string, status *coreTypes.Status) StatusContainer {
	ret := StatusContainer{}
	// EXISTING_CODE
	ret.Chain = chain
	ret.Status = *status
	// TODO: This is a hack. We need to get the version from the core
	ret.Version = version.LibraryVersion
	ret.Caches = status.Caches
	ret.LastUpdate = time.Now()
	// EXISTING_CODE
	return ret
}

func (s *StatusContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *StatusContainer) NeedsUpdate(force bool) bool {
	elapsed := time.Now().After(s.LastUpdate.Add(time.Minute * 2))
	if force || elapsed {
		s.LastUpdate = time.Now()
		return true
	}
	return false
}

func (s *StatusContainer) ShallowCopy() Containerer {
	ret := &StatusContainer{
		Status:     s.Status.ShallowCopy(),
		NItems:     s.NItems,
		NFolders:   s.NFolders,
		NFiles:     s.NFiles,
		NBytes:     s.NBytes,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	ret.Status.Chain = s.Status.Chain
	return ret
}

func (s *StatusContainer) Summarize() {
	// EXISTING_CODE
	s.NItems = len(s.Caches)
	for _, cache := range s.Caches {
		s.NFolders += int(cache.NFolders)
		s.NFiles += int(cache.NFiles)
		s.NBytes += int(cache.SizeInBytes)
	}
	// EXISTING_CODE
}

func StatusX() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
