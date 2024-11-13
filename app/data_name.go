// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var namesMutex sync.Mutex
var namesChain = "mainnet"

// EXISTING_CODE

var nameLock atomic.Uint32

func (a *App) loadNames(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !nameLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer nameLock.CompareAndSwap(1, 0)

	if !a.names.NeedsUpdate(a.forceName()) {
		return nil
	}

	opts := sdk.NamesOptions{
		Globals: a.getGlobals(),
	}
	// EXISTING_CODE
	names.ClearCustomNames()
	opts.All = true
	// EXISTING_CODE
	opts.Verbose = true

	if names, meta, err := opts.NamesList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (names == nil) || (len(names) == 0) {
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
		a.names = types.NewNameContainer(opts.Chain, names)
		// EXISTING_CODE
		a.namesMap = make(map[base.Address]coreTypes.Name, len(names))
		for _, name := range names {
			a.namesMap[name.Address] = name
		}
		// EXISTING_CODE
		a.emitInfoMsg("Loaded names", "")
	}

	return nil
}

func (a *App) forceName() (force bool) {
	// EXISTING_CODE
	tm := file.MustGetLatestFileTime(coreConfig.MustGetPathToChainConfig(namesChain))
	ret := tm.Unix()
	force = ret > a.names.LastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (a *App) ModifyName(modData *ModifyData) error {
	if !nameLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer nameLock.CompareAndSwap(1, 0)

	op := modData.Operation
	newName := coreTypes.Name{
		Address:  modData.Address,
		Name:     modData.Value,
		IsCustom: true,
		Source:   "TrueBlocks Browse",
		Tags:     "99-User-Defined",
	}
	if existing, ok := a.namesMap[modData.Address]; ok {
		if existing.IsCustom {
			// We preserve the tags if it's already customized
			newName.Tags = existing.Tags
		}
	}

	cd := crud.CrudFromName(newName)
	opts := sdk.NamesOptions{
		Globals: a.getGlobals(),
	}
	opts.Globals.Chain = namesChain

	if _, _, err := opts.ModifyName(crud.OpFromString(op), cd); err != nil {
		a.emitErrorMsg(err, nil)
		return err
	}

	newArray := []coreTypes.Name{}
	for _, name := range a.names.Items {
		if name.Address == modData.Address {
			switch crud.OpFromString(op) {
			case crud.Update:
				name = newName
			default:
				if name.IsCustom {
					// we can only delete if it's custom already
					switch crud.OpFromString(op) {
					case crud.Delete:
						name.Deleted = true
					case crud.Undelete:
						name.Deleted = false
					case crud.Remove:
						continue
					}
				}
			}
			namesMutex.Lock()
			a.namesMap[modData.Address] = name
			namesMutex.Unlock()
		}
		newArray = append(newArray, name)
	}
	namesMutex.Lock()
	a.names.Items = newArray
	namesMutex.Unlock()

	return nil
}

// EXISTING_CODE
