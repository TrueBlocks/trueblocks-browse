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
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var namesMutex sync.Mutex
var namesChain = "mainnet"

// EXISTING_CODE

var namesLock atomic.Uint32

func (a *App) loadNames(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadNames", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !namesLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer namesLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.names.NeedsUpdate() {
		return nil
	}
	updater := a.names.Updater
	defer func() {
		a.names.Updater = updater
	}()
	logger.InfoBY("Updating needed for names...")

	opts := sdk.NamesOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	opts.All = true
	// EXISTING_CODE
	if items, meta, err := opts.NamesList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no names found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		namesMutex.Lock()
		defer namesMutex.Unlock()
		// EXISTING_CODE
		a.meta = *meta
		a.names = types.NewNameContainer(opts.Chain, items)
		// EXISTING_CODE
		a.namesMap = make(map[base.Address]types.Name, len(items))
		for _, name := range items {
			a.namesMap[name.Address] = name
		}
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "names")
	}

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
