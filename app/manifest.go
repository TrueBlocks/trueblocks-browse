package app

import (
	"fmt"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetManifest(first, pageSize int) types.ManifestContainer {
	first = base.Max(0, base.Min(first, len(a.manifest.Items)-1))
	last := base.Min(len(a.manifest.Items), first+pageSize)
	copy := a.manifest.ShallowCopy()
	copy.Items = a.manifest.Items[first:last]
	return copy
}

func (a *App) GetManifestCnt() int {
	return len(a.manifest.Items)
}

func (a *App) loadManifest(wg *sync.WaitGroup) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	opts := sdk.ChunksOptions{}
	if manifests, _, err := opts.ChunksManifest(); err != nil {
		return err
	} else if (manifests == nil) || (len(manifests) == 0) {
		return fmt.Errorf("no manifest found")
	} else {
		if len(a.manifest.Items) == len(manifests[0].Chunks) {
			return nil
		}
		a.manifest = types.NewSummaryManifest(manifests[0])
		sort.Slice(a.manifest.Items, func(i, j int) bool {
			return a.manifest.Items[i].Range > a.manifest.Items[j].Range
		})
	}
	return nil
}
