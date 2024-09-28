package utils

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func MustGetLatestFileTime(path string) time.Time {
	if info, err := file.GetNewestInDirectory(path); err != nil {
		logger.Error("error getting latest file time", "error", err)
		return time.Now()
	} else {
		return info.ModTime()
	}
}
