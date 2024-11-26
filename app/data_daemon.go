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
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var daemonsLock atomic.Uint32

func (a *App) loadDaemons(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadDaemons", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !daemonsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer daemonsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.daemons.NeedsUpdate() {
		return nil
	}
	updater := a.daemons.Updater
	defer func() {
		a.daemons.Updater = updater
	}()
	logger.InfoBY("Updating needed for daemons...")

	opts := DaemonsOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	opts.Chain = a.getChain()
	// EXISTING_CODE
	if items, meta, err := opts.DaemonsList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		// expected outcome
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.daemons = types.NewDaemonContainer(opts.Chain, items)
		// EXISTING_CODE
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "daemons")
	}

	return nil
}

// EXISTING_CODE
type DaemonsOptions struct {
	Globals sdk.Globals
	Chain   string
}

func (opts *DaemonsOptions) DaemonsList() ([]types.Daemon, *coreTypes.MetaData, error) {
	meta, err := sdk.GetMetaData(namesChain)
	// TODO: We've been called to check status, do wizard checks here
	return []types.Daemon{}, meta, err
}

// EXISTING_CODE
