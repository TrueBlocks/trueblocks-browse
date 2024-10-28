package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var indexLock atomic.Uint32

func (a *App) IndexPage(first, pageSize int) *types.IndexContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.indexes.Items)-1))
	last := base.Min(len(a.indexes.Items), first+pageSize)
	copy, _ := a.indexes.ShallowCopy().(*types.IndexContainer)
	copy.Items = a.indexes.Items[first:last]
	return copy
}

func (a *App) loadIndexes(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !indexLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer indexLock.CompareAndSwap(1, 0)

	if !a.indexes.NeedsUpdate(a.forceIndex()) {
		return nil
	}

	opts := sdk.IndexesOptions{
		Globals: a.toGlobals(),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	opts.Verbose = true

	if indexes, meta, err := opts.IndexesList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (indexes == nil) || (len(indexes) == 0) {
		err = fmt.Errorf("no indexes found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.indexes = types.NewIndexContainer(opts.Chain, indexes)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := sdk.SortIndexes(a.indexes.Items, a.indexes.Sorts); err != nil {
			messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
				String1: err.Error(),
			})
		}
		a.indexes.Summarize()
		messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded indexes"})
	}
	return nil
}

func (a *App) forceIndex() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
