package app

import (
	"fmt"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// Find: NewViews
func (a *App) GetIndex(first, pageSize int) types.IndexSummary {
	first = base.Max(0, base.Min(first, len(a.index.Chunks)-1))
	last := base.Min(len(a.index.Chunks), first+pageSize)
	copy := a.index.ShallowCopy()
	copy.Chunks = a.index.Chunks[first:last]
	return copy
}

func (a *App) GetIndexCnt() int {
	return len(a.index.Chunks)
}

func (a *App) loadIndex() error {
	var err error
	opts := sdk.ChunksOptions{}
	if a.index.Chunks, _, err = opts.ChunksStats(); err != nil {
		return err
	} else if (a.index.Chunks == nil) || (len(a.index.Chunks) == 0) {
		return fmt.Errorf("no index chunks found")
	} else {
		// reverse order
		sort.Slice(a.index.Chunks, func(i, j int) bool {
			return a.index.Chunks[i].Range > a.index.Chunks[j].Range
		})
		a.index.Summarize()
	}
	return nil
}
