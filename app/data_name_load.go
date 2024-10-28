package app

// EXISTING_CODE
import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

func (a *App) nameLoadMe() {
	name := types.NewNameContainer(a.Chain, []coreTypes.Name{})
	fmt.Println("Loaded: ", name.String())
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
