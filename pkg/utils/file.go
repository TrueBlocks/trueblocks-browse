package utils

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func MustGetLatestFileTime(folders ...string) time.Time {
	if len(folders) == 0 {
		return time.Now()
	}

	getTimeOfNewestFile := func(folder string) time.Time {
		if info, err := file.GetNewestInDirectory(folder); err != nil || info == nil {
			if info == nil {
				logger.Warn("latest file time skipped", "error", err)
			} else {
				logger.Error("error getting latest file time:", err)
			}
			return time.Now()
		} else {
			return info.ModTime()
		}
	}

	var newestTime time.Time
	for _, folder := range folders {
		fileTime := getTimeOfNewestFile(folder)
		if fileTime.After(newestTime) {
			newestTime = fileTime
		}
	}

	return newestTime
}
