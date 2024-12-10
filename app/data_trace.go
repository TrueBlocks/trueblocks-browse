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

var tracesLock atomic.Uint32

func (a *App) loadTraces(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadTraces", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !tracesLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer tracesLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.traces.NeedsUpdate() {
		return nil
	}
	updater := a.traces.Updater
	defer func() {
		a.traces.Updater = updater
	}()
	logger.InfoBY("Updating traces...")

	if items, meta, err := a.pullTraces(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no traces found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.traces = types.NewTraceContainer(a.getChain(), items, a.GetLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.traces.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "traces")
	}

	return nil
}

func (a *App) pullTraces() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
