package app

// EXISTING_CODE
import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
)

// EXISTING_CODE

func (a *App) configLoadMe() {
	config := types.NewConfigContainer(a.Chain, []configTypes.Config{})
	fmt.Println("Loaded: ", config.String())
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
