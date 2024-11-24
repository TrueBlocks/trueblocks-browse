package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/editors"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

func (a *App) LoadName(addr string) editors.Name {
	if name, ok := a.namesMap[base.HexToAddress(addr)]; ok {
		logger.Info("Found name for ", name.Address.Hex())
		return editors.CoreToName(name)
	} else {
		logger.Info("Could not find name for ", name.Address.Hex())
		return editors.CoreToName(types.Name{
			Name:     "Unnamed",
			Address:  base.HexToAddress("0x0"),
			Tags:     "99-User-Defined",
			Source:   "TrueBlocks Browse",
			Symbol:   "",
			Decimals: 18,
			Deleted:  false,
		})
	}
}

// func (a *App) SaveName(name editors.Name) error {
// 	logger.Info("Setting name", name.String())
// 	if name.Name == "Error" {
// 		return fmt.Errorf("not implemented")
// 	}
// 	time.Sleep(200 * time.Millisecond)
// 	return nil
// }
