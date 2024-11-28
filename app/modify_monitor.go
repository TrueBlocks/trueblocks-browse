package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) ModifyMonitor(modData *ModifyData) error {
	if !monitorsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer monitorsLock.CompareAndSwap(1, 0)

	op := crud.OpFromString(modData.Operation)
	opts := sdk.MonitorsOptions{
		Addrs:    []string{modData.Address.Hex()},
		Delete:   op == crud.Delete,
		Undelete: op == crud.Undelete,
		Remove:   op == crud.Remove,
		Globals: sdk.Globals{
			Chain:   a.getChain(),
			Verbose: false,
		},
	}

	if _, _, err := opts.Monitors(); err != nil {
		a.emitErrorMsg(err, nil)
		return err
	}

	newArray := []types.Monitor{}
	for _, mon := range a.monitors.Items {
		if mon.Address == modData.Address {
			switch op {
			case crud.Delete:
				mon.Deleted = true
			case crud.Undelete:
				mon.Deleted = false
			case crud.Remove:
				continue
			}
		}
		newArray = append(newArray, mon)
	}
	monitorMutex.Lock()
	defer monitorMutex.Unlock()

	a.monitors.Items = newArray
	return nil
}
