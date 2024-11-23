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

	if !a.session.NeedsUpdate() {
		return nil
	}
	logger.InfoBY("Updating needed for session...")

	opts := sdk.SessionOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	if session, meta, err := opts.SessionList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (session == nil) || (len(session) == 0) {
		err = fmt.Errorf("no session found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// Oddly, the session is always up to date, so we save it...
		ss := a.session.Session
		// EXISTING_CODE
		a.meta = *meta
		a.session = types.NewSessionContainer(opts.Chain, &session[0])
		// EXISTING_CODE
		// ... and put it back
		a.session.Session = ss
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "session")
	}

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
