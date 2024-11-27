// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"io"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var statusLock atomic.Uint32

func (a *App) loadStatus(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadStatus", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !statusLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer statusLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.status.NeedsUpdate() {
		return nil
	}
	updater := a.status.Updater
	defer func() {
		a.status.Updater = updater
	}()
	logger.InfoBY("Updating status...")

	if items, meta, err := a.pullStatus(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no status found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.status = types.NewStatusContainer(a.getChain(), items)
		// EXISTING_CODE
		// TODO: Use the core's sorting mechanism (see SortChunk Stats for example)
		sort.Slice(a.status.Caches, func(i, j int) bool {
			return a.status.Caches[i].SizeInBytes > a.status.Caches[j].SizeInBytes
		})
		// EXISTING_CODE
		if err := sdk.SortStatus(a.status.Items, a.status.Sorts); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "status")
	}

	return nil
}

func (a *App) pullStatus() (items []types.Status, meta *types.Meta, err error) {
	// EXISTING_CODE
	w := logger.GetLoggerWriter()
	logger.SetLoggerWriter(io.Discard)
	defer logger.SetLoggerWriter(w)
	opts := sdk.StatusOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	return opts.StatusList()
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
