package app

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) IndexPage(first, pageSize int) *types.IndexContainer {
	first = base.Max(0, base.Min(first, len(a.index.Items)-1))
	last := base.Min(len(a.index.Items), first+pageSize)
	copy, _ := a.index.ShallowCopy().(*types.IndexContainer)
	copy.Items = a.index.Items[first:last]
	return copy
}

var indexLock atomic.Uint32

func (a *App) loadIndex(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !indexLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer indexLock.CompareAndSwap(1, 0)

	if !a.index.NeedsUpdate() {
		return nil
	}

	chain := a.globals.Chain
	opts := sdk.ChunksOptions{
		Globals: sdk.Globals{
			Verbose: true,
			Chain:   chain,
		},
	}

	if chunks, meta, err := opts.ChunksStats(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (chunks == nil) || (len(chunks) == 0) {
		err = fmt.Errorf("no index chunks found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		a.index = types.NewIndexContainer(chain, chunks)
		if err := sdk.SortChunkStats(a.index.Items, a.index.Sorts); err != nil {
			messages.SendError(a.ctx, err)
		}
		a.index.Summarize()
		messages.SendInfo(a.ctx, "Loaded indexes")
	}
	return nil
}
