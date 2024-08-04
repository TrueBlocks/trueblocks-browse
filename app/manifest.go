package app

import (
	"fmt"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetManifest(first, pageSize int) types.ManifestSummary {
	manifest := types.NewManifestEx(a.manifest)
	first = base.Max(0, base.Min(first, len(manifest.Chunks)-1))
	last := base.Min(len(manifest.Chunks), first+pageSize)
	manifest.Chunks = manifest.Chunks[first:last]
	return manifest
}

func (a *App) GetManifestCnt() int {
	return len(a.manifest.Chunks)
}

func (a *App) loadManifest() error {
	opts := sdk.ChunksOptions{}
	if manifestArray, _, err := opts.ChunksManifest(); err != nil {
		return err
	} else if (manifestArray == nil) || (len(manifestArray) == 0) {
		return fmt.Errorf("no manifest found")
	} else {
		sort.Slice(manifestArray[0].Chunks, func(i, j int) bool {
			return manifestArray[0].Chunks[i].Range > manifestArray[0].Chunks[j].Range
		})
		a.manifest = manifestArray[0]
	}
	return nil
}
