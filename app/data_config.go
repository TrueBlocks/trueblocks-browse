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
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var configLock atomic.Uint32

func (a *App) ConfigPage(first, pageSize int) *types.ConfigContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.configs.Items)-1))
	last := base.Min(len(a.configs.Items), first+pageSize)
	copy, _ := a.configs.ShallowCopy().(*types.ConfigContainer)
	copy.Items = a.configs.Items[first:last]
	return copy
}

func (a *App) loadConfigs(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !configLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer configLock.CompareAndSwap(1, 0)

	if !a.configs.NeedsUpdate(a.forceConfig()) {
		return nil
	}

	opts := sdk.ConfigsOptions{
		Globals: a.toGlobals(),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	opts.Verbose = true

	if configs, meta, err := opts.ConfigsList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (configs == nil) || (len(configs) == 0) {
		err = fmt.Errorf("no configs found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.configs = types.NewConfigContainer(opts.Chain, configs)
		// EXISTING_CODE
		// EXISTING_CODE
		a.configs.Summarize()
		messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded configs"})
	}
	return nil
}

func (a *App) forceConfig() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
