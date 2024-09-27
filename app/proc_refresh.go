package app

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

var freshenLock atomic.Uint32
var freshenMutex sync.Mutex

// Refresh when the app starts and then later by the daemons to instruct the backend and
// by extension the frontend to update. We protect against updating too fast... Note
// that this routine is called as a goroutine.
func (a *App) Refresh(skipable bool, which ...string) {
	if !a.isConfigured() {
		return
	}

	// Skip this update we're actively upgrading
	if skipable {
		if !freshenLock.CompareAndSwap(0, 1) {
			// logger.Info(colors.Red, "Skipping update", colors.Off)
			return
		}
		defer freshenLock.CompareAndSwap(1, 0)
	}

	freshenMutex.Lock()
	defer freshenMutex.Unlock()

	if !a.ScraperController.IsRunning() {
		logger.Info(colors.Green, "Freshening...", colors.Off)
	}

	// We always load names first since we need them everywhere
	err := a.loadNames(nil, nil)
	if err != nil {
		// we report the error, but proceed anyway
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(
			err,
		))
	}

	// We want to update the route we last used first if there is one...
	if len(which) > 0 {
		switch which[0] {
		case "/abis":
			err = a.loadAbis(nil, nil)
		case "/manifest":
			err = a.loadManifest(nil, nil)
		case "/monitors":
			err = a.loadMonitors(nil, nil)
		case "/index":
			err = a.loadIndex(nil, nil)
		case "/status":
			err = a.loadStatus(nil, nil)
		}
		if err != nil {
			// we report the error, but proceed anyway
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(
				err,
			))
		}
	}
	// a.loadPortfolio(nil, nil)

	// And then update everything else in the fullness of time
	wg := sync.WaitGroup{}
	errorChan := make(chan error, 5) // Buffered channel to hold up to 5 errors (one from each goroutine)

	wg.Add(6)
	go a.loadNames(&wg, errorChan)
	go a.loadAbis(&wg, errorChan)
	go a.loadManifest(&wg, errorChan)
	go a.loadMonitors(&wg, errorChan)
	go a.loadIndex(&wg, errorChan)
	go a.loadStatus(&wg, errorChan)

	go func() {
		wg.Wait()
		close(errorChan) // Close the channel after all goroutines are done
	}()

	var errors []error
	for err := range errorChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	a.loadPortfolio(nil, nil)

	if len(errors) > 0 {
		// Handle errors, e.g., wait 1/2 second between each error message
		for _, err := range errors {
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(
				err,
			))
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		messages.Send(a.ctx, messages.Daemon, messages.NewDaemonMsg(
			a.FreshenController.Name,
			"Freshening...",
			a.FreshenController.Color,
		))
	}
}
