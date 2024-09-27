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

func (a *App) StatusPage(first, pageSize int) types.StatusContainer {
	first = base.Max(0, base.Min(first, len(a.status.Items)-1))
	last := base.Min(len(a.status.Items), first+pageSize)
	copy := a.status.ShallowCopy()
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

	// silence progress reporting for a second...
	w := logger.GetLoggerWriter()
	logger.SetLoggerWriter(io.Discard)
	defer logger.SetLoggerWriter(w)

	opts := sdk.StatusOptions{
		Globals: a.globals,
	}

	messages.SendInfo(a.ctx, "Freshening status")
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
		a.status = types.NewStatusContainer(statusArray[0])
		sort.Slice(a.status.Items, func(i, j int) bool {
			return a.status.Items[i].SizeInBytes > a.status.Items[j].SizeInBytes
		})
		a.status.Summarize()
		messages.SendInfo(a.ctx, "Finished loading status")
	}
	return nil
}
