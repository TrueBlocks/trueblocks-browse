// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
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

	if !a.settings.NeedsUpdate() {
		return nil
	}

	// EXISTING_CODE
	_ = errorChan // delint
	if path, err := utils.GetConfigFn("", "trueBlocks.toml"); err != nil {
		a.emitErrorMsg(err, nil)
	} else {
		if err := coreConfig.ReadToml(path, &a.config.Config); err != nil {
			a.emitErrorMsg(err, nil)
		}
	}

	props := types.SettingsProps{
		Status:  &a.status,
		Config:  &a.config,
		Session: &a.session,
	}
	a.settings = types.NewSettingsContainer(a.getChain(), &props)
	// EXISTING_CODE
	// EXISTING_CODE
	// do not remove
	// EXISTING_CODE
	// EXISTING_CODE
	// do not remove
	// EXISTING_CODE
	a.emitLoadingMsg(messages.Loaded, "settings")

	return nil
}

// EXISTING_CODE
// EXISTING_CODE
