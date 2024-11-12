// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var abisChain = "mainnet"

// EXISTING_CODE

var abiLock atomic.Uint32

func (a *App) loadAbis(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !abiLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer abiLock.CompareAndSwap(1, 0)

	if !a.abis.NeedsUpdate(a.forceAbi()) {
		return nil
	}

	opts := sdk.AbisOptions{
		Globals: a.getGlobals(),
	}
	// EXISTING_CODE
	opts.Cache = true
	opts.Chain = abisChain
	// EXISTING_CODE
	opts.Verbose = true

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
		a.abis.Summarize()
		a.emitInfoMsg("Loaded abis", "")
	}

	return nil
}

func (a *App) ModifyAbi(modData *ModifyData) error {
	opts := sdk.AbisOptions{
		Addrs:   []string{modData.Address.Hex()},
		Globals: a.getGlobals(),
	}
	opts.Globals.Decache = true

	if _, _, err := opts.Abis(); err != nil {
		a.emitAddressErrorMsg(err, modData.Address)
		return err
	} else {
		newAbis := make([]coreTypes.Abi, 0, len(a.abis.Items))
		for _, abi := range a.abis.Items {
			if abi.Address == modData.Address {
				a.abis.NItems--
				a.abis.NEvents -= abi.NEvents
				a.abis.NFunctions -= abi.NFunctions
				continue
			}
			newAbis = append(newAbis, abi)
		}
		a.abis.LastUpdate = 0
		a.abis.Items = newAbis
		a.emitInfoMsg("ModifyAbi delete", modData.Address.Hex())
		return nil
	}
}

func (a *App) forceAbi() (force bool) {
	// EXISTING_CODE
	force = a.forceName()
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
