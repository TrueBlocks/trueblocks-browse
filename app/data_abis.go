package app

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) AbiPage(first, pageSize int) *types.AbiContainer {
	first = base.Max(0, base.Min(first, len(a.abis.Items)-1))
	last := base.Min(len(a.abis.Items), first+pageSize)
	copy, _ := a.abis.ShallowCopy().(*types.AbiContainer)
	copy.Items = a.abis.Items[first:last]
	return copy
}

var abisChain = "mainnet"
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

	if !a.abis.NeedsUpdate(a.nameChange()) {
		return nil
	}

	opts := sdk.AbisOptions{
		Globals: a.globals,
	}
	opts.Globals.Chain = abisChain

	opts.Globals.Verbose = true
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
		a.meta = *meta
		a.abis = types.NewAbiContainer(abisChain, abis)
		if err := sdk.SortAbis(a.abis.Items, a.abis.Sorts); err != nil {
			messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
				String1: err.Error(),
			})
		}
		a.abis.Summarize()
		messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded abis"})
	}
	return nil
}

func (a *App) ModifyAbi(modData *ModifyData) error {
	opts := sdk.AbisOptions{
		Addrs:   []string{modData.Address.Hex()},
		Globals: a.globals,
	}
	opts.Globals.Decache = true

	if _, _, err := opts.Abis(); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
			Address: modData.Address,
		})
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
		a.abis.LastUpdate = time.Time{}
		a.abis.Items = newAbis
		msg := fmt.Sprintf("ModifyAbi delete: %s", modData.Address.Hex())
		messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: msg})
		return nil
	}
}
