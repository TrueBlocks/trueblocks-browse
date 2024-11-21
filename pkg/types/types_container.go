package types

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
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

func DebugInts(label string, lastUp, up walk.Updater) {
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
		logger.InfoBG("NeedsUpdate:", label, "lastUpdate", lastUp.String(), "latest", up.String())
	}
}

