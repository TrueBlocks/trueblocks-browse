package app

import (
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) Logger(msg []string) {
	logger.Info(msg)
}

var isTesting bool

func init() {
	isTesting = os.Getenv("TB_TEST_MODE") == "true"
}
