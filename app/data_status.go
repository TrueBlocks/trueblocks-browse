package app

import (
	"fmt"
	"io"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) StatusPage(first, pageSize int) *types.StatusContainer {
	first = base.Max(0, base.Min(first, len(a.status.Caches)-1))
	last := base.Min(len(a.status.Caches), first+pageSize)
	copy, _ := a.status.ShallowCopy().(*types.StatusContainer)
	copy.Caches = a.status.Caches[first:last]
	return copy
}

var statusLock atomic.Uint32

func (a *App) loadStatus(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !statusLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer statusLock.CompareAndSwap(1, 0)

	if !a.status.NeedsUpdate(false) {
		return nil
	}

	// silence progress reporting for a second...
	w := logger.GetLoggerWriter()
	logger.SetLoggerWriter(io.Discard)
	defer logger.SetLoggerWriter(w)

	chain := a.globals.Chain
	opts := sdk.StatusOptions{
		Chains:  true,
		Globals: a.globals,
	}

	if statusArray, meta, err := opts.StatusAll(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (statusArray == nil) || (len(statusArray) == 0) {
		err = fmt.Errorf("no status found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		a.status = types.NewStatusContainer(chain, statusArray[0])
		// TODO: Use the core's sorting mechanism (see SortChunkStats for example)
		sort.Slice(a.status.Caches, func(i, j int) bool {
			return a.status.Caches[i].SizeInBytes > a.status.Caches[j].SizeInBytes
		})
		a.status.Summarize()
		logger.SetLoggerWriter(w)
		messages.EmitInfo(a.ctx, "Loaded status")
	}
	return nil
}

func (a *App) GetChainInfo(chain string) coreTypes.Chain {
	for _, ch := range a.status.Chains {
		if ch.Chain == chain {
			return ch
		}
	}
	return coreTypes.Chain{}
}
