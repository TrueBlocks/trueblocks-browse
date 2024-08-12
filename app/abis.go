package app

import (
	"fmt"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// Find: NewViews
func (a *App) GetAbis(first, pageSize int) types.SummaryAbis {
	first = base.Max(0, base.Min(first, len(a.abis.Items)-1))
	last := base.Min(len(a.abis.Items), first+pageSize)
	copy := a.abis.ShallowCopy()
	copy.Items = a.abis.Items[first:last]
	return copy
}

func (a *App) GetAbisCnt() int {
	return len(a.abis.Items)
}

func (a *App) loadAbis(wg *sync.WaitGroup) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

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
		if len(a.abis.Items) == len(abis) {
			return nil
		}
		a.abis = types.SummaryAbis{}
		a.abis.Items = append(a.abis.Items, abis...)
		a.abis.Summarize()
	}
	return nil
}
