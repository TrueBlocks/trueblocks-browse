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

var publishLock atomic.Uint32

func (a *App) loadPublish(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadPublish", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !publishLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer publishLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.publish.NeedsUpdate() {
		return nil
	}
	updater := a.publish.Updater
	defer func() {
		a.publish.Updater = updater
	}()
	logger.InfoBY("Updating publish...")

	if items, meta, err := a.pullPublishs(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no publish found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.publish = types.NewPublishContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.publish.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "publish")
	}

	return nil
}

func (a *App) pullPublishs() (items []types.CacheItem, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
