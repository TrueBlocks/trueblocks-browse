package app

import (
	"fmt"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/version"
)

func (a *App) GetStatus(first, pageSize int) types.SummaryStatus {
	first = base.Max(0, base.Min(first, len(a.status.Items)-1))
	last := base.Min(len(a.status.Items), first+pageSize)
	copy := a.status.ShallowCopy()
	copy.Items = a.status.Items[first:last]
	return copy
}

func (a *App) GetStatusCnt() int {
	return len(a.status.Items)
}

func (a *App) loadStatus() error {
	opts := sdk.StatusOptions{}
	if statusArray, _, err := opts.StatusAll(); err != nil {
		return err
	} else if (statusArray == nil) || (len(statusArray) == 0) {
		return fmt.Errorf("no status found")
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
