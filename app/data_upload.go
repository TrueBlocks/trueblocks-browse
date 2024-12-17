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
)

// EXISTING_CODE

var uploadsLock atomic.Uint32

func (a *App) loadUploads(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadUploads", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !uploadsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer uploadsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.uploads.NeedsUpdate() {
		return nil
	}
	updater := a.uploads.Updater
	defer func() {
		a.uploads.Updater = updater
	}()
	logger.InfoBY("Updating uploads...")

	if items, meta, err := a.pullUploads(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no uploads found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.uploads = types.NewUploadContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.uploads.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "uploads")
	}

	return nil
}

func (a *App) pullUploads() (items []types.CacheItem, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
