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
	logger.InfoBY("Updating names...")

	if items, meta, err := a.pullNames(); err != nil {
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
		a.namesMap = make(map[base.Address]types.Name, len(items))
		for _, name := range items {
			a.namesMap[name.Address] = name
		}
		// EXISTING_CODE
		a.meta = *meta
		a.names = types.NewNameContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "names")
	}

	return nil
}

func (a *App) pullNames() (items []types.Name, meta *types.Meta, err error) {
	// EXISTING_CODE
	opts := sdk.NamesOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	opts.All = true
	return opts.NamesList()
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
