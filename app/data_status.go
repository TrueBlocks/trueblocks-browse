package app

import (
	"fmt"
	"io"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) StatusPage(first, pageSize int) *types.StatusContainer {
	first = base.Max(0, base.Min(first, len(a.status.Items)-1))
	last := base.Min(len(a.status.Items), first+pageSize)
	copy, _ := a.status.ShallowCopy().(*types.StatusContainer)
	copy.Items = a.status.Items[first:last]
	return copy
}

func (a *App) loadStatus(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	if !a.status.NeedsUpdate() {
		return nil
	}

	// silence progress reporting for a second...
	w := logger.GetLoggerWriter()
	logger.SetLoggerWriter(io.Discard)
	defer logger.SetLoggerWriter(w)

	chain := a.globals.Chain
	opts := sdk.StatusOptions{
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
		sort.Slice(a.status.Items, func(i, j int) bool {
			return a.status.Items[i].SizeInBytes > a.status.Items[j].SizeInBytes
		})
		a.status.Summarize()
		logger.SetLoggerWriter(w)
		messages.SendInfo(a.ctx, "Loaded status")
	}
	return nil
}
