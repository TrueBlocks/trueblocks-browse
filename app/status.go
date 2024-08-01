package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetStatus(first, pageSize int) types.StatusEx {
	status := types.NewStatusEx(a.status)
	first = base.Max(0, base.Min(first, len(status.Caches)-1))
	last := base.Min(len(status.Caches), first+pageSize)
	status.Caches = status.Caches[first:last]
	return status
}

func (a *App) GetStatusCnt() int {
	return len(a.status.Caches)
}

func (a *App) loadStatus() error {
	opts := sdk.StatusOptions{}
	if statusArray, _, err := opts.StatusAll(); err != nil {
		return err
	} else if (statusArray == nil) || (len(statusArray) == 0) {
		return fmt.Errorf("no status found")
	} else {
		a.status = statusArray[0]
	}
	return nil
}
