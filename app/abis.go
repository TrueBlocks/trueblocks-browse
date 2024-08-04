package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// Find: NewViews
func (a *App) GetAbisPage(first, pageSize int) types.AbiSummary {
	copy := a.abis
	first = base.Max(0, base.Min(first, len(copy.Files)-1))
	last := base.Min(len(copy.Files), first+pageSize)
	copy.Files = copy.Files[first:last]
	return copy
}

func (a *App) GetAbisCnt() int {
	return len(a.abis.Files)
}

func (a *App) loadAbis() error {
	opts := sdk.AbisOptions{
		Globals: sdk.Globals{
			Verbose: true,
		},
	}
	if abis, _, err := opts.AbisList(); err != nil {
		return err
	} else if (abis == nil) || (len(abis) == 0) {
		return fmt.Errorf("no status found")
	} else {
		a.abis.Files = append(a.abis.Files, abis...)
		a.abis.Summarize()
	}
	return nil
}
