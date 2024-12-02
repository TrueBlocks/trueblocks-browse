// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var manifestsLock atomic.Uint32

func (a *App) loadManifests(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadManifests", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !manifestsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer manifestsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.manifests.NeedsUpdate() {
		return nil
	}
	updater := a.manifests.Updater
	defer func() {
		a.manifests.Updater = updater
	}()
	logger.InfoBY("Updating manifests...")

	opts := sdk.ChunksOptions{
		Globals: sdk.Globals{
			Chain:   a.getChain(),
			Verbose: true,
		},
	}
	// EXISTING_CODE
	// EXISTING_CODE
	if manifests, meta, err := opts.ChunksManifest(); err != nil {
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
		if err := sdk.SortChunkRecords(a.manifests.Items, a.manifests.Sorts); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "manifests")
	}

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
