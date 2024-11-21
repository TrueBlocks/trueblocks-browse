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
)

var abisChain = "mainnet"

// EXISTING_CODE

var abiLock atomic.Uint32

func (a *App) loadAbis(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadAbis", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !abiLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer abiLock.CompareAndSwap(1, 0)

	if !a.abis.NeedsUpdate() {
		return nil
	}

	opts := sdk.AbisOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	opts.Cache = true
	opts.Chain = abisChain
	// EXISTING_CODE
	if abis, meta, err := opts.AbisList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (abis == nil) || (len(abis) == 0) {
		err = fmt.Errorf("no abis found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.abis = types.NewAbiContainer(opts.Chain, abis)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := sdk.SortAbis(a.abis.Items, a.abis.Sorts); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitInfoMsg("Loaded abis", "")
	}

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
