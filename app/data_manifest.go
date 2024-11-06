// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var manifestLock atomic.Uint32

func (a *App) ManifestPage(first, pageSize int) *types.ManifestContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.manifests.Items)-1))
	last := base.Min(len(a.manifests.Items), first+pageSize)
	copy, _ := a.manifests.ShallowCopy().(*types.ManifestContainer)
	copy.Items = a.manifests.Items[first:last]
	return copy
}

func (a *App) loadManifests(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !manifestLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer manifestLock.CompareAndSwap(1, 0)

	if !a.manifests.NeedsUpdate(a.forceManifest()) {
		return nil
	}

	opts := sdk.ManifestsOptions{
		Globals: a.getGlobals(),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	opts.Verbose = true

	if manifests, meta, err := opts.ManifestsList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (manifests == nil) || (len(manifests) == 0) {
		err = fmt.Errorf("no manifests found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.manifests = types.NewManifestContainer(opts.Chain, manifests)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := sdk.SortManifests(a.manifests.Items, a.manifests.Sorts); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.manifests.Summarize()
		a.emitInfoMsg("Loaded manifests", "")
	}

	return nil
}

func (a *App) forceManifest() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
