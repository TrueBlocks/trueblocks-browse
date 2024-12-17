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

var receiptsLock atomic.Uint32

func (a *App) loadReceipts(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadReceipts", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !receiptsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer receiptsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.receipts.NeedsUpdate() {
		return nil
	}
	updater := a.receipts.Updater
	defer func() {
		a.receipts.Updater = updater
	}()
	logger.InfoBY("Updating receipts...")

	if items, meta, err := a.pullReceipts(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no receipts found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.receipts = types.NewReceiptContainer(a.getChain(), items, a.getLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.receipts.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "receipts")
	}

	return nil
}

func (a *App) pullReceipts() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
