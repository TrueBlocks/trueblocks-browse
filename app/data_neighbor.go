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
)

// EXISTING_CODE

var neighborsLock atomic.Uint32

func (a *App) loadNeighbors(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadNeighbors", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !neighborsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer neighborsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.neighbors.NeedsUpdate() {
		return nil
	}
	updater := a.neighbors.Updater
	defer func() {
		a.neighbors.Updater = updater
	}()
	logger.InfoBY("Updating neighbors...")

	if items, meta, err := a.pullNeighbors(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no neighbors found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.neighbors = types.NewNeighborContainer(a.getChain(), items, a.getLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.neighbors.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "neighbors")
	}

	return nil
}

func (a *App) pullNeighbors() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
