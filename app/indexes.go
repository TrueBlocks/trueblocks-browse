package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// Find: NewViews
func (a *App) GetIndexesPage(first, pageSize int) []coreTypes.ChunkStats {
	first = base.Max(0, base.Min(first, len(a.indexes)-1))
	last := base.Min(len(a.indexes), first+pageSize)
	return a.indexes[first:last]
}

func (a *App) GetIndexesCnt() int {
	return len(a.indexes)
}

func (a *App) loadIndexes() error {
	opts := sdk.ChunksOptions{}
	if indexArray, _, err := opts.ChunksStats(); err != nil {
		return err
	} else if (indexArray == nil) || (len(indexArray) == 0) {
		return fmt.Errorf("no status found")
	} else {
		a.indexes = indexArray
	}
	return nil
}
