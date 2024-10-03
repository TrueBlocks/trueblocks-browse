package app

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var nameMutex sync.Mutex

// Find: NewViews
func (a *App) NamePage(first, pageSize int) *types.NameContainer {
	nameMutex.Lock()
	defer nameMutex.Unlock()

	first = base.Max(0, base.Min(first, len(a.names.Names)-1))
	last := base.Min(len(a.names.Names), first+pageSize)
	copy, _ := a.names.ShallowCopy().(*types.NameContainer)
	copy.Names = a.names.Names[first:last]
	return copy
}

var namesChain = "mainnet"
var namesLock atomic.Uint32

func (a *App) loadNames(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !namesLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer namesLock.CompareAndSwap(1, 0)

	if !a.names.NeedsUpdate() {
		return nil
	}

	names.ClearCustomNames()

	parts := coreTypes.Regular | coreTypes.Custom | coreTypes.Prefund | coreTypes.Baddress
	if namesMap, err := names.LoadNamesMap(namesChain, parts, nil); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (namesMap == nil) || (len(namesMap) == 0) {
		err = fmt.Errorf("no names found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// a.meta = *meta
		nameMutex.Lock()
		defer nameMutex.Unlock()

		a.names = types.NewNameContainer(namesChain, namesMap)
		for _, name := range a.names.NamesMap {
			a.names.Names = append(a.names.Names, name)
		}
		sort.Slice(a.names.Names, func(i, j int) bool {
			return compare(a.names.Names[i], a.names.Names[j])
		})
		a.names.Summarize()
		messages.SendInfo(a.ctx, "Loaded names")
	}
	return nil
}

func compare(nameI, nameJ coreTypes.Name) bool {
	ti := nameI.Parts
	if ti == coreTypes.Regular {
		ti = 7
	}
	tj := nameJ.Parts
	if tj == coreTypes.Regular {
		tj = 7
	}
	if ti == tj {
		if nameI.Tags == nameJ.Tags {
			return nameI.Address.Hex() < nameJ.Address.Hex()
		}
		return nameI.Tags < nameJ.Tags
	}
	return ti < tj
}

func (a *App) ModifyName(op string, address base.Address) error {
	if !a.isConfigured() {
		return nil
	}

	opFromString := func(op string) crud.NameOperation {
		m := map[string]crud.NameOperation{
			"delete":   crud.Delete,
			"undelete": crud.Undelete,
			"remove":   crud.Remove,
		}
		return m[op]
	}

	cd := crud.CrudFromAddress(address)
	messages.SendInfo(a.ctx, fmt.Sprintf("%s-%v", opFromString(op), *cd))

	if _, ok := a.names.NamesMap[address]; ok {
		opts := sdk.NamesOptions{
			Globals: a.globals,
		}
		opts.Globals.Chain = namesChain
		if _, _, err := opts.ModifyName(opFromString(op), cd); err != nil {
			messages.SendError(a.ctx, err)
			return err
		}

		nameMutex.Lock()
		defer nameMutex.Unlock()
		newArray := []coreTypes.Name{}
		for _, n := range a.names.Names {
			if n.IsCustom && n.Address == address {
				switch opFromString(op) {
				case crud.Delete:
					n.Deleted = true
				case crud.Undelete:
					n.Deleted = false
				case crud.Remove:
					continue
				}
				a.names.NamesMap[address] = n
			}
			newArray = append(newArray, n)
		}
		a.names.Names = newArray
	}

	return nil
}
