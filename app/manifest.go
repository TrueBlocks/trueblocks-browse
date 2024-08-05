package app

import (
	"fmt"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetManifest(first, pageSize int) types.ManifestSummary {
	first = base.Max(0, base.Min(first, len(a.manifest.Chunks)-1))
	last := base.Min(len(a.manifest.Chunks), first+pageSize)
	copy := a.manifest.ShallowCopy()
	copy.Chunks = a.manifest.Chunks[first:last]
	return copy
}

func (a *App) GetManifestCnt() int {
	return len(a.manifest.Chunks)
}

func (a *App) loadManifest() error {
	opts := sdk.ChunksOptions{}
	if manifests, _, err := opts.ChunksManifest(); err != nil {
		return err
	} else if (manifests == nil) || (len(manifests) == 0) {
		return fmt.Errorf("no manifest found")
	} else {
		a.manifest = types.NewManifestEx(manifests[0])
		sort.Slice(a.manifest.Chunks, func(i, j int) bool {
			return a.manifest.Chunks[i].Range > a.manifest.Chunks[j].Range
		})
	}
	return nil
}
