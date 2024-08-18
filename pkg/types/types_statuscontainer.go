package types

import (
	"time"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type StatusContainer struct {
	coreTypes.Status `json:",inline"`
	Items            []coreTypes.CacheItem `json:"items"`
	NItems           int                   `json:"nItems"`
	LatestUpdate     string                `json:"latestUpdate"`
	NFolders         int                   `json:"nFolders"`
	NFiles           int                   `json:"nFiles"`
	NBytes           int                   `json:"nBytes"`
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
	s.NItems = len(s.Items)
	s.LatestUpdate = time.Now().Format(time.RFC3339)
	for _, cache := range s.Items {
		s.NFolders += int(cache.NFolders)
		s.NFiles += int(cache.NFiles)
		s.NBytes += int(cache.SizeInBytes)
	}
}
