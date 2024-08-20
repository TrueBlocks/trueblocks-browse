package app

import (
	"fmt"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// Find: NewViews
func (a *App) IndexPage(first, pageSize int) types.IndexContainer {
	first = base.Max(0, base.Min(first, len(a.index.Items)-1))
	last := base.Min(len(a.index.Items), first+pageSize)
	copy := a.index.ShallowCopy()
	copy.Items = a.index.Items[first:last]
	return copy
}

func (a *App) loadIndex(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	opts := sdk.ChunksOptions{}
	if chunks, meta, err := opts.ChunksStats(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (chunks == nil) || (len(chunks) == 0) {
		err = fmt.Errorf("no index chunks found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		if len(a.index.Items) == len(chunks) {
			return nil
		}
		a.index = types.IndexContainer{Items: chunks}
		sort.Slice(a.index.Items, func(i, j int) bool {
			// reverse order
			return a.index.Items[i].Range > a.index.Items[j].Range
		})
		a.index.Summarize()
	}
	return nil
}
