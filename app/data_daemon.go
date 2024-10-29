// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

// EXISTING_CODE

var daemonLock atomic.Uint32

func (a *App) DaemonPage(first, pageSize int) *types.DaemonContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	copy := &types.DaemonContainer{} // a.daemons.ShallowCopy().(*types.DaemonContainer)
	return copy
}

func (a *App) loadDaemons(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !daemonLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer daemonLock.CompareAndSwap(1, 0)

	// if !a.daemons.NeedsUpdate(a.forceDaemon()) {
	// 	return nil
	// }

	// opts := sdk.DaemonsOptions{
	// 	Globals: a.toGlobals(),
	// }
	// // EXISTING_CODE
	// // EXISTING_CODE
	// opts.Verbose = true

	// if daemons, meta, err := opts.DaemonsList(); err != nil {
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else if (daemons == nil) || (len(daemons) == 0) {
	// 	err = fmt.Errorf("no daemons found")
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else {
	// 	// EXISTING_CODE
	// 	// EXISTING_CODE
	// 	a.meta = *meta
	// 	a.daemons = types.NewDaemonContainer(opts.Chain, &daemons[0])
	// 	// EXISTING_CODE
	// 	// EXISTING_CODE
	// 	a.daemons.Summarize()
	// 	messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded daemons"})
	// }
	return nil
}

func (a *App) forceDaemon() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
