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
	render := true
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
	default:
		// do nothing
	}

	if render {
		now := time.Now().Unix()
		logger.InfoW(label, "lastUpdate", now-lastUpdate, "latest", now-latest)
	}
}

// var dateFmt string = "15:04:05"
