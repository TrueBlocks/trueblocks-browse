package app

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var nameMutex sync.Mutex
var namesChain = "mainnet"
var namesLock atomic.Uint32

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

	if !a.names.NeedsUpdate(false) {
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
		messages.EmitInfo(a.ctx, "Loaded names")
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

func (a *App) ModifyName(modData *ModifyData) error {
	if !namesLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer namesLock.CompareAndSwap(1, 0)

	opFromString := func(op string) crud.Operation {
		m := map[string]crud.Operation{
			"update":   crud.Update,
			"delete":   crud.Delete,
			"undelete": crud.Undelete,
			"remove":   crud.Remove,
		}
		return m[op]
	}

	op := modData.Operation
	newName := coreTypes.Name{
		Address:  modData.Address,
		Name:     modData.Value,
		IsCustom: true,
		Source:   "TrueBlocks Browse",
		Tags:     "99-User-Defined",
	}
	if existing, ok := a.names.NamesMap[modData.Address]; ok {
		if existing.IsCustom {
			// We preserve the tags if it's already customized
			newName.Tags = existing.Tags
		}
	}

	cd := crud.CrudFromName(newName)
	opts := sdk.NamesOptions{
		Globals: a.globals,
	}
	opts.Globals.Chain = namesChain

	if _, _, err := opts.ModifyName(opFromString(op), cd); err != nil {
		messages.EmitError(a.ctx, err)
		return err
	}

	newArray := []coreTypes.Name{}
	for _, name := range a.names.Names {
		if name.Address == modData.Address {
			switch opFromString(op) {
			case crud.Update:
				name = newName
			default:
				if name.IsCustom {
					// we can only delete if it's custom already
					switch opFromString(op) {
					case crud.Delete:
						name.Deleted = true
					case crud.Undelete:
						name.Deleted = false
					case crud.Remove:
						continue
					}
				}
			}
			nameMutex.Lock()
			a.names.NamesMap[modData.Address] = name
			nameMutex.Unlock()
		}
		newArray = append(newArray, name)
	}
	nameMutex.Lock()
	a.names.Names = newArray
	nameMutex.Unlock()

	return nil
}

func (a *App) nameChange() bool {
	latest := utils.MustGetLatestFileTime(config.MustGetPathToChainConfig(namesChain))
	return latest != a.names.LastUpdate
}
