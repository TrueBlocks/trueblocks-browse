package app

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"

func debugMsg(v ...any) {
	logger.Info(v...)
}
