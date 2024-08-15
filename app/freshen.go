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

// Refresh when the app starts and then later by the daemons to instruct the backend and
// by extension the frontend to update. We protect against updating too fast... Note
// that this routine is called as a goroutine.
func (a *App) Refresh(which ...string) {
	// Skip this update we're actively upgrading
	if !freshenLock.CompareAndSwap(0, 1) {
		// logger.Info(colors.Red, "Skipping update", colors.Off)
		return
	}
	logger.Info(colors.Green, "Freshening...", colors.Off)
	defer freshenLock.CompareAndSwap(1, 0)

	// Function to let the front end know that something freshened
	notify :=
		func(msg messages.Message, msgStr string) {
			messages.Send(a.ctx, msg, messages.NewDaemonMsg(
				a.FreshenController.Name,
				msgStr,
				a.FreshenController.Color,
			))
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
		}
		if err != nil {
			// we report the error, but proceed anyway
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(
				err,
			))
		} else {
			notify(messages.Daemon, "Freshing...")
		}
	}

	// And then update everything else in the fullness of time
	wg := sync.WaitGroup{}
	errorChan := make(chan error, 5) // Buffered channel to hold up to 5 errors (one from each goroutine)

	wg.Add(5)
	go a.loadNames(&wg, errorChan)
	go a.loadAbis(&wg, errorChan)
	go a.loadManifest(&wg, errorChan)
	go a.loadMonitors(&wg, errorChan)
	go a.loadIndex(&wg, errorChan)

	wg.Wait()
	close(errorChan) // Close the channel after all goroutines are done

	var errors []error
	for err := range errorChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		// Handle errors, e.g., wait 1/2 second between each error message
		for _, err := range errors {
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(
				err,
			))
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		notify(messages.Daemon, "Freshing...")
	}
}
