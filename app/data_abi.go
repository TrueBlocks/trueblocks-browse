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

var abisLock atomic.Uint32

func (a *App) loadAbis(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadAbis", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !abisLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer abisLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.abis.NeedsUpdate() {
		return nil
	}
	updater := a.abis.Updater
	defer func() {
		a.abis.Updater = updater
	}()
	logger.InfoBY("Updating abis...")

	if items, meta, err := a.pullAbis(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no abis found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.abis = types.NewAbiContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.abis.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "abis")
	}

	return nil
}

func (a *App) pullAbis() (items []types.Abi, meta *types.Meta, err error) {
	// EXISTING_CODE
	opts := sdk.AbisOptions{
		Globals: sdk.Globals{
			Chain:   namesChain,
			Verbose: true,
			Cache:   true,
		},
	}
	return opts.AbisList()
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
