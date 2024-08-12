package types

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually this will get put back into Core.

type StatusContainer struct {
	coreTypes.Status `json:",inline"`
	Items            []coreTypes.CacheItem `json:"items"`
	LatestUpdate     string                `json:"latestUpdate"`
	NFolders         uint64                `json:"nFolders"`
	NFiles           uint64                `json:"nFiles"`
	NBytes           int64                 `json:"nBytes"`
}

func (s *StatusContainer) ShallowCopy() StatusContainer {
	return StatusContainer{
		Status: s.Status.ShallowCopy(),
	}
}

func (s *StatusContainer) Summarize() {
	for _, cache := range s.Items {
		s.NFolders += cache.NFolders
		s.NFiles += cache.NFiles
		s.NBytes += cache.SizeInBytes
	}
}
