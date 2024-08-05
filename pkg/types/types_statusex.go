package types

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually this will get put back into Core.

type SummaryStatus struct {
	coreTypes.Status `json:",inline"`
	LatestUpdate     string `json:"latestUpdate"`
	NFolders         uint64 `json:"nFolders"`
	NFiles           uint64 `json:"nFiles"`
	NBytes           int64  `json:"nBytes"`
}

func (s *SummaryStatus) ShallowCopy() SummaryStatus {
	return SummaryStatus{
		Status: s.Status.ShallowCopy(),
	}
}

func (s *SummaryStatus) Summarize() {
	for _, cache := range s.Caches {
		s.NFolders += cache.NFolders
		s.NFiles += cache.NFiles
		s.NBytes += cache.SizeInBytes
	}
}
