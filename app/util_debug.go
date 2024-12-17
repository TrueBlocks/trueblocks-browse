package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) Logger(msg []string) {
	logger.InfoBY(msg)
}
