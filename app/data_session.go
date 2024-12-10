// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var sessionLock atomic.Uint32

func (a *App) loadSession(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadSession", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !sessionLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer sessionLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.session.NeedsUpdate() {
		return nil
	}
	updater := a.session.Updater
	defer func() {
		a.session.Updater = updater
	}()
	logger.InfoBY("Updating session...")

	if items, meta, err := a.pullSessions(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		// this outcome is okay
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.session = types.NewSessionContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.session.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "session")
	}

	return nil
}

func (a *App) pullSessions() (items []types.Session, meta *types.Meta, err error) {
	// EXISTING_CODE
	meta, err = sdk.GetMetaData(namesChain)
	items = []types.Session{}
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
