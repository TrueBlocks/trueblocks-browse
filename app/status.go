package app

import (
	"fmt"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/version"
)

func (a *App) GetStatus(first, pageSize int) types.StatusContainer {
	first = base.Max(0, base.Min(first, len(a.status.Items)-1))
	last := base.Min(len(a.status.Items), first+pageSize)
	copy := a.status.ShallowCopy()
	copy.Items = a.status.Items[first:last]
	return copy
}

func (a *App) GetStatusCnt() int {
	return len(a.status.Items)
}

func (a *App) loadStatus(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	opts := sdk.StatusOptions{}
	if statusArray, _, err := opts.StatusAll(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (statusArray == nil) || (len(statusArray) == 0) {
		err = fmt.Errorf("no status found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.status.Status = statusArray[0]
		// TODO: This is a hack. We need to get the version from the core
		a.status.Version = version.LibraryVersion
		a.status.Items = a.status.Caches
		sort.Slice(a.status.Items, func(i, j int) bool {
			return a.status.Items[i].SizeInBytes > a.status.Items[j].SizeInBytes
		})
		a.status.Summarize()
	}
	return nil
}
