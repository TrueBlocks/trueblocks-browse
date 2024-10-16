package app

import (
	"fmt"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/editors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) LoadName(addr string) editors.Name {
	name := coreTypes.Name{
		Name:     "My Name",
		Address:  base.HexToAddress("0xf503017d7baf7fbc0fff7492b751025c6a78179b"),
		Tags:     "99-User-Defined",
		Source:   "TrueBlocks",
		Symbol:   "",
		Decimals: 18,
		Deleted:  false,
	}

	return editors.CoreToName(name)
}

func (a *App) SaveName(name editors.Name) error {
	logger.Info("Setting name", name.String())
	if name.Name == "Error" {
		return fmt.Errorf("not implemented")
	}
	time.Sleep(200 * time.Millisecond)
	return nil
}
