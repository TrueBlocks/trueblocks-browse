package app

import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

// EXISTING_CODE

var sessionLock atomic.Uint32

func (a *App) SessionPage(first, pageSize int) *types.SessionContainer {
	// EXISTING_CODE
	_ = first
	_ = pageSize
	// EXISTING_CODE

	copy, _ := a.sessions.ShallowCopy().(*types.SessionContainer)
	return copy
}

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

	if !a.sessions.NeedsUpdate(a.forceSession()) {
		return nil
	}

	_ = errorChan
	// opts := sdk.SessionsOptions{
	// 	Globals: a.toGlobals(),
	// }
	// // EXISTING_CODE
	// // EXISTING_CODE
	// opts.Verbose = true

	// if sessions, meta, err := opts.SessionsList(); err != nil {
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else if (sessions == nil) || (len(sessions) == 0) {
	// 	err = fmt.Errorf("no sessions found")
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else {
	// 	// EXISTING_CODE
	// 	// EXISTING_CODE
	// 	a.meta = *meta
	// 	a.sessions = types.NewSessionContainer(opts.Chain, &sessions[0])
	// 	// EXISTING_CODE
	// 	// EXISTING_CODE
	// 	a.sessions.Summarize()
	// 	messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded sessions"})
	// }
	return nil
}

func (a *App) forceSession() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
