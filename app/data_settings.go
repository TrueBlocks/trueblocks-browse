// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var settingsLock atomic.Uint32

func (a *App) loadSettings(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadSettings", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !settingsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer settingsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.settings.NeedsUpdate() {
		return nil
	}
	updater := a.settings.Updater
	defer func() {
		a.settings.Updater = updater
	}()
	logger.InfoBY("Updating settings...")

	if items, meta, err := a.pullSettings(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		// this outcome is okay
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.settings = types.NewSettingsContainer(a.getChain(), items)
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.settings.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "settings")
	}

	return nil
}

func (a *App) pullSettings() (items []types.CacheItem, meta *types.Meta, err error) {
	// EXISTING_CODE
	if meta, err = sdk.GetMetaData(a.getChain()); err != nil {
		return nil, nil, err
	} else {
		if err = a.loadSession(nil, nil); err != nil {
			return nil, nil, err
		}
		if err = a.loadStatus(nil, nil); err != nil {
			return nil, nil, err
		}
		if err = a.loadConfig(nil, nil); err != nil {
			return nil, nil, err
		}
		a.settings.Status = a.status
		a.settings.Config = a.config
		a.settings.Session = a.session
		return a.status.Items, meta, nil
	}
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
