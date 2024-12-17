package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) LoadName(addr string) types.Name {
	if name, ok := a.namesMap[base.HexToAddress(addr)]; ok {
		logger.Info("Found name for ", name.Address.Hex())
		return name
	} else {
		logger.Info("Could not find name for ", name.Address.Hex())
		return types.Name{
			Name:     "Unnamed",
			Address:  base.HexToAddress(addr),
			Tags:     "99-User-Defined",
			Source:   "TrueBlocks Browse",
			Symbol:   "",
			Decimals: 18,
			Deleted:  false,
		}
	}
}

func (a *App) SaveName(name types.Name) error {
	logger.InfoBC("SaveName", name)
	if name.Name == "Error" {
		return fmt.Errorf("not implemented")
	}
	return nil
}
