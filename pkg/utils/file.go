package utils

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func MustGetLatestFileTime(path string) time.Time {
	if info, err := file.GetNewestInDirectory(path); err != nil || info == nil {
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
