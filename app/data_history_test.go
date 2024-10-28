package app

import (
	"testing"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func TestHistoryLoadMe(t *testing.T) {
	app := NewApp()
	app.loadHistory(base.ZeroAddr, nil, nil)
}
