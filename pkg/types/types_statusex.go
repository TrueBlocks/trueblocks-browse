package types

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually this will get put back into Core.

type StatusEx struct {
	coreTypes.Status `json:",inline"`
	LatestUpdate     string `json:"latestUpdate"`
	NFolders         uint64 `json:"nFolders"`
	NFiles           uint64 `json:"nFiles"`
	NBytes           int64  `json:"nBytes"`
}

func NewStatusEx(status coreTypes.Status) StatusEx {
	ret := StatusEx{
		Status: status,
	}

	for _, cache := range status.Caches {
		ret.NFolders += cache.NFolders
		ret.NFiles += cache.NFiles
		ret.NBytes += cache.SizeInBytes
	}

	return ret
}
