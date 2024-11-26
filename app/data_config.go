// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var configLock atomic.Uint32

func (a *App) loadConfig(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadConfig", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !configLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer configLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.config.NeedsUpdate() {
		return nil
	}
	updater := a.config.Updater
	defer func() {
		a.config.Updater = updater
	}()
	logger.InfoBY("Updating needed for config...")

	opts := sdk.ConfigOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	if config, meta, err := opts.ConfigList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (config == nil) || (len(config) == 0) {
		err = fmt.Errorf("no config found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.config = types.NewConfigContainer(opts.Chain, config)
		// EXISTING_CODE
		sort.Slice(a.config.Items, func(i, j int) bool {
			return a.config.Items[i].ChainId < a.config.Items[j].ChainId
		})
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "config")
	}

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
