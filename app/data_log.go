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

var logsLock atomic.Uint32

func (a *App) loadLogs(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadLogs", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !logsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer logsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.logs.NeedsUpdate() {
		return nil
	}
	updater := a.logs.Updater
	defer func() {
		a.logs.Updater = updater
	}()
	logger.InfoBY("Updating logs...")

	if items, meta, err := a.pullLogs(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no logs found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.logs = types.NewLogContainer(a.getChain(), items, a.getLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.logs.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "logs")
	}

	return nil
}

func (a *App) pullLogs() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
