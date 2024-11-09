// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// EXISTING_CODE

var sessionLock atomic.Uint32

func (a *App) loadSessions(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !sessionLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer sessionLock.CompareAndSwap(1, 0)

	if !a.session.NeedsUpdate(a.forceSession()) {
		return nil
	}

	opts := sdk.SessionsOptions{
		Globals: a.getGlobals(),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	opts.Verbose = true

	if sessions, meta, err := opts.SessionsList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (sessions == nil) || (len(sessions) == 0) {
		err = fmt.Errorf("no sessions found")
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
		a.session = types.NewSessionContainer(opts.Chain, &sessions[0])
		// EXISTING_CODE
		// ... and put it back
		a.session.Session = ss
		// EXISTING_CODE
		a.session.Summarize()
		a.emitInfoMsg("Loaded sessions", "")
	}
	return nil
}

func (a *App) forceSession() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (a *App) saveSession() {
	if !isTesting {
		a.session.Window.X, a.session.Window.Y = runtime.WindowGetPosition(a.ctx)
		a.session.Window.Width, a.session.Window.Height = runtime.WindowGetSize(a.ctx)
		a.session.Window.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	}
	_ = a.session.Save()
}

// EXISTING_CODE
