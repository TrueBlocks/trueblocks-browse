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
	logger.InfoBY("Updating config...")

	if items, meta, err := a.pullConfigs(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no config found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.config = types.NewConfigContainer(a.getChain(), items)
		// EXISTING_CODE
		sort.Slice(a.config.Items, func(i, j int) bool {
			return a.config.Items[i].ChainId < a.config.Items[j].ChainId
		})
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "config")
	}

	return nil
}

func (a *App) pullConfigs() (items []types.Config, meta *types.Meta, err error) {
	// EXISTING_CODE
	opts := sdk.ConfigOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	return opts.ConfigList()
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
