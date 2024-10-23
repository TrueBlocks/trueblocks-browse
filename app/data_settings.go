package app

import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
)

func (a *App) SettingsPage(first, pageSize int) *types.SettingsGroup {
	return &a.settings
}

var settingsLock atomic.Uint32

func (a *App) loadSettings(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !settingsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer settingsLock.CompareAndSwap(1, 0)

	if !a.settings.NeedsUpdate(false) {
		return nil
	}

	if path, err := utils.GetConfigFn("", "trueBlocks.toml"); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
		})
	} else {
		if err := coreConfig.ReadToml(path, &a.cfg); err != nil {
			messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
				String1: err.Error(),
			})
		}
	}
	a.settings = types.NewSettingsGroup(&a.status.Status, &a.cfg, &a.session)
	a.settings.Summarize()

	return nil
}
