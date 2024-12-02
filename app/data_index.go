// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var indexLock atomic.Uint32

func (a *App) loadIndexes(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadIndexes", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !indexLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer indexLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.indexes.NeedsUpdate() {
		return nil
	}
	updater := a.indexes.Updater
	defer func() {
		a.indexes.Updater = updater
	}()
	logger.InfoBY("Updating needed for indexes...")

	opts := sdk.ChunksOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	if indexes, meta, err := opts.ChunksStats(); err != nil {
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
		if err := sdk.SortChunkStats(a.indexes.Items, a.indexes.Sorts); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "indexes")
	}

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
