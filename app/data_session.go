package app

// EXISTING_CODE
import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

func (a *App) sessionLoadMe() {
	session := types.NewSessionContainer(a.Chain, &config.Session{})
	fmt.Println("Loaded: ", session.String())
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
