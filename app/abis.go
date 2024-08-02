package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// Find: NewViews
func (a *App) GetAbisPage(first, pageSize int) []coreTypes.AbiFile {
	first = base.Max(0, base.Min(first, len(a.abis)-1))
	last := base.Min(len(a.abis), first+pageSize)
	return a.abis[first:last]
}

func (a *App) GetAbisCnt() int {
	return len(a.abis)
}

func (a *App) loadAbis() error {
	opts := sdk.AbisOptions{}
	if abisArray, _, err := opts.AbisList(); err != nil {
		return err
	} else if (abisArray == nil) || (len(abisArray) == 0) {
		return fmt.Errorf("no status found")
	} else {
		a.abis = abisArray
	}
	return nil
}
