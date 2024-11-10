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
	GetItems() interface{}
	SetItems(items interface{})
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
