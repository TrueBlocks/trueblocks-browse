package app

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

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
		Globals: a.getGlobals(false /* verbose */),
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