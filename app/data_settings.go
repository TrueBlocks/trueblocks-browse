// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
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

	_ = errorChan // delint
	if path, err := utils.GetConfigFn("", "trueBlocks.toml"); err != nil {
		a.emitErrorMsg(err, nil)
	} else {
		if err := coreConfig.ReadToml(path, &a.config.Config); err != nil {
			a.emitErrorMsg(err, nil)
		}
	}

	a.settings = types.NewSettingsGroup(&a.status.Status, &a.config.Config, &a.session.Session)
	a.settings.Summarize()

	return nil
}

/*
func (a *App) S tatusPage(first, pageSize int) *types.S tatusContainer {
	first = base.Max(0, base.Min(first, len(a.status.Caches)-1))
	last := base.Min(len(a.status.Caches), first+pageSize)
	copy, _ := a.status.ShallowCopy().(*types.S tatusContainer)
	copy.Caches = a.status.Caches[first:last]
	return copy
}

var statusLock atomic.Uint32

func (a *App) l oadStatus(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !statusLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer statusLock.CompareAndSwap(1, 0)

	if !a.status.NeedsUpdate(false) {
		return nil
	}

	// silence progress reporting for a second...
	w := logger.GetLoggerWriter()
	logger.SetLoggerWriter(io.Discard)
	defer logger.SetLoggerWriter(w)

	chain := a.globals.Chain
	opts := sdk.StatusOptions{
		Chains:  true,
		Globals: a.toGlobals(),
	}

	if statusArray, meta, err := opts.StatusAll(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (statusArray == nil) || (len(statusArray) == 0) {
		err = fmt.Errorf("no status found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		a.meta = *meta
		a.status = types.NewS tatusContainer(chain, &statusArray[0])
		// TODO: Use the core's sorting mechanism (see SortChunkStats for example)
		sort.Slice(a.status.Caches, func(i, j int) bool {
			return a.status.Caches[i].SizeInBytes > a.status.Caches[j].SizeInBytes
		})
		a.status.Summarize()
		logger.SetLoggerWriter(w)
		a.emitInfoMsg("Loaded status")
	}
	return nil
}
*/
