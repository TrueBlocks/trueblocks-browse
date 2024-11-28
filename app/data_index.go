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

var indexesLock atomic.Uint32

func (a *App) loadIndexes(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadIndexes", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !indexesLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer indexesLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.indexes.NeedsUpdate() {
		return nil
	}
	updater := a.indexes.Updater
	defer func() {
		a.indexes.Updater = updater
	}()
	logger.InfoBY("Updating indexes...")

	if items, meta, err := a.pullIndexes(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no indexes found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.indexes = types.NewIndexContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.indexes.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "indexes")
	}

	return nil
}

func (a *App) pullIndexes() (items []types.ChunkStats, meta *types.Meta, err error) {
	// EXISTING_CODE
	opts := sdk.ChunksOptions{
		Globals: sdk.Globals{
			Chain:   a.getChain(),
			Verbose: true,
		},
	}
	return opts.ChunksStats()
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
