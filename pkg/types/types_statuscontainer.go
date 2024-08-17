package types

import (
	"time"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type StatusContainer struct {
	coreTypes.Status `json:",inline"`
	Items            []coreTypes.CacheItem `json:"items"`
	NItems           uint64                `json:"nItems"`
	LatestUpdate     string                `json:"latestUpdate"`
	NFolders         uint64                `json:"nFolders"`
	NFiles           uint64                `json:"nFiles"`
	NBytes           int64                 `json:"nBytes"`
}

func (s *StatusContainer) ShallowCopy() StatusContainer {
	return StatusContainer{
		Status:       s.Status.ShallowCopy(),
		NItems:       s.NItems,
		LatestUpdate: s.LatestUpdate,
		NFolders:     s.NFolders,
		NFiles:       s.NFiles,
		NBytes:       s.NBytes,
	}
}

func (s *StatusContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	s.LatestUpdate = time.Now().Format(time.RFC3339)
	for _, cache := range s.Items {
		s.NFolders += cache.NFolders
		s.NFiles += cache.NFiles
		s.NBytes += cache.SizeInBytes
	}
}
