package app

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) ModifyAbi(modData *ModifyData) error {
	opts := sdk.AbisOptions{
		Addrs:   []string{modData.Address.Hex()},
		Globals: a.getGlobals(false /* verbose */),
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
		a.abis.Updater.Reset()
		a.abis.Items = newAbis
		a.emitInfoMsg("ModifyAbi delete", modData.Address.Hex())
		return nil
	}
}
