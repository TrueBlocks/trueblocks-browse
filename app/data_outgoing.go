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

var outgoingLock atomic.Uint32

func (a *App) loadOutgoing(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadOutgoing", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !outgoingLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer outgoingLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.outgoing.NeedsUpdate() {
		return nil
	}
	updater := a.outgoing.Updater
	defer func() {
		a.outgoing.Updater = updater
	}()
	logger.InfoBY("Updating outgoing...")

	if items, meta, err := a.pullOutgoing(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no outgoing found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.outgoing = types.NewOutgoingContainer(a.getChain(), items, a.GetLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.outgoing.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "outgoing")
	}

	return nil
}

func (a *App) pullOutgoing() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
