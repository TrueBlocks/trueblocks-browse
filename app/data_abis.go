package app

import (
	"fmt"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) AbiPage(first, pageSize int) *types.AbiContainer {
	first = base.Max(0, base.Min(first, len(a.abis.Items)-1))
	last := base.Min(len(a.abis.Items), first+pageSize)
	copy, _ := a.abis.ShallowCopy().(*types.AbiContainer)
	copy.Items = a.abis.Items[first:last]
	return copy
}

var abisChain = "mainnet"

func (a *App) loadAbis(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	if !a.abis.NeedsUpdate() {
		return nil
	}

	opts := sdk.AbisOptions{
		Globals: a.globals,
	}
	opts.Globals.Chain = abisChain

	opts.Globals.Verbose = true
	if abis, meta, err := opts.AbisList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (abis == nil) || (len(abis) == 0) {
		err = fmt.Errorf("no abis found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		a.abis = types.NewAbiContainer(abisChain, abis)
		if err := sdk.SortAbis(a.abis.Items, a.abis.Sorts); err != nil {
			messages.SendError(a.ctx, err)
		}
		a.abis.Summarize()
		messages.SendInfo(a.ctx, "Loaded abis")
	}
	return nil
}
