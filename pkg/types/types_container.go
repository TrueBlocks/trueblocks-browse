package types

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	CollateAndFilter(theMap *FilterMap) interface{}
	NeedsUpdate() bool
	GetItems() interface{}
	SetItems(items interface{})
}

type Containerers []Containerer

var debugging = true

func DebugInts(label string, lastUpdate, latest int64) {
	if !debugging {
		return
	}

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
		logger.InfoBG("NeedsUpdate:", label, "lastUpdate", now-lastUpdate, "latest", now-latest)
	}
}

// var dateFmt string = "15:04:05"
