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

var incomingLock atomic.Uint32

func (a *App) loadIncoming(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadIncoming", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !incomingLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer incomingLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.incoming.NeedsUpdate() {
		return nil
	}
	updater := a.incoming.Updater
	defer func() {
		a.incoming.Updater = updater
	}()
	logger.InfoBY("Updating incoming...")

	if items, meta, err := a.pullIncoming(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no incoming found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.incoming = types.NewIncomingContainer(a.getChain(), items, a.GetLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.incoming.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "incoming")
	}

	return nil
}

func (a *App) pullIncoming() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE