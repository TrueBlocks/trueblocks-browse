package app

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// Find: NewViews
func (a *App) ManifestPage(first, pageSize int) *types.ManifestContainer {
	first = base.Max(0, base.Min(first, len(a.manifest.Items)-1))
	last := base.Min(len(a.manifest.Items), first+pageSize)
	copy, _ := a.manifest.ShallowCopy().(*types.ManifestContainer)
	copy.Items = a.manifest.Items[first:last]
	return copy
}

var manifestLock atomic.Uint32

func (a *App) loadManifest(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !manifestLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer manifestLock.CompareAndSwap(1, 0)

	if !a.manifest.NeedsUpdate(false) {
		return nil
	}

	chain := a.Chain
	opts := sdk.ChunksOptions{
		Globals: sdk.Globals{
			Verbose: true,
			Chain:   chain,
		},
	}

	if manifests, meta, err := opts.ChunksManifest(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (manifests == nil) || (len(manifests) == 0) {
		err = fmt.Errorf("no manifest found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		a.manifest = types.NewManifestContainer(chain, manifests[0].Chunks)
		a.manifest.Version = manifests[0].Version
		a.manifest.Specification = string(manifests[0].Specification)
		// TODO: Use sorting mechanism from core (see SortChunkStats for example)
		sort.Slice(a.manifest.Items, func(i, j int) bool {
			return a.manifest.Items[i].Range > a.manifest.Items[j].Range
		})
		a.manifest.Summarize()
		messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded manifest"})
	}
	return nil
}
