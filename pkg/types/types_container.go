package types

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	Summarize()
	NeedsUpdate(force bool) bool
}

type Containerers []Containerer

func DebugInts(label string, lastUpdate, latest int64) {
	now := time.Now().Unix()
	logger.InfoW(label, "lastUpdate", now-lastUpdate, "latest", now-latest)
}

// var dateFmt string = "15:04:05"
