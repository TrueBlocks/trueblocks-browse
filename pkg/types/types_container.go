package types

import (
	"fmt"
	"os"

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

var debugging = false

func init() {
	debugging = os.Getenv("TB_DEBUG_INTS") == "true"
}

func DebugInts(label string, lastUp, up walk.Updater) {
	if !debugging {
		return
	}

	render := true
	switch label {
	case "abis":
	case "config":
	case "daemons":
	case "history":
	case "indexes":
	case "manifests":
	case "monitors":
	case "names":
	case "project":
	case "session":
	case "status":
	case "wizard":
	case "settings":
	default:
		logger.Fatal("Invalid label in DebugInts")
		// do nothing
	}

	if render {
		logger.InfoBY(fmt.Sprintf("DB2: % 5.5s: % 11d, % 11d, % 11d", label, lastUp.LastTs, up.LastTs, up.LastTs-lastUp.LastTs))
	}
}
