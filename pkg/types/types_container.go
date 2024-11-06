package types

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	Summarize()
	NeedsUpdate(meta *coreTypes.MetaData, force bool) bool
}

type Containerers []Containerer

func DebugInts(label string, lastUpdate, latest int64) {
	switch label {
	case "abi":
	case "config":
	case "daemon":
	case "history":
	case "index":
	case "manifest":
	case "monitor":
	case "name":
	case "project":
	case "session":
	case "status":
	case "wizard":
		now := time.Now().Unix()
		logger.InfoW(label, "lastUpdate", now-lastUpdate, "latest", now-latest)
	default:
		// do nothing
	}
}

// var dateFmt string = "15:04:05"
