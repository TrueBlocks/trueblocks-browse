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

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

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

	if !a.status.NeedsUpdate(a.forceStatus()) {
		return nil
	}

	opts := sdk.StatusOptions{
		Globals: a.getGlobals(),
	}
	// EXISTING_CODE
	w := logger.GetLoggerWriter()
	logger.SetLoggerWriter(io.Discard)
	defer logger.SetLoggerWriter(w)
	// EXISTING_CODE
	opts.Verbose = true

	if status, meta, err := opts.StatusList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (status == nil) || (len(status) == 0) {
		err = fmt.Errorf("no status found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.status = types.NewStatusContainer(opts.Chain, &status[0])
		// EXISTING_CODE
		// TODO: Use the core's sorting mechanism (see SortChunkStats for example)
		sort.Slice(a.status.Caches, func(i, j int) bool {
			return a.status.Caches[i].SizeInBytes > a.status.Caches[j].SizeInBytes
		})
		logger.SetLoggerWriter(w)
		// EXISTING_CODE
		a.status.Summarize()
		a.emitInfoMsg("Loaded status", "")
	}

	return nil
}

func (a *App) forceStatus() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
