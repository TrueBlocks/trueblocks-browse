package app

import (
	"fmt"
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
func (a *App) Refresh() error {
	if !a.IsConfigured() {
		return fmt.Errorf("App not configured")
	}

	if !freshenLock.CompareAndSwap(0, 1) {
		return nil // it's okay to skip a refresh if one is already in progress
	}
	defer freshenLock.CompareAndSwap(1, 0)

	freshenMutex.Lock()
	defer freshenMutex.Unlock()

	if !a.ScraperController.IsRunning() {
		logger.Info(colors.Green, "Freshening...", colors.Off)
	}

	// We always load names first since we need them everywhere
	err := a.loadNames(nil, nil)
	if err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
		})
	}

	// And then update everything else in the fullness of time
	wg := sync.WaitGroup{}
	errorChan := make(chan error, 5) // Buffered channel to hold up to 5 errors (one from each goroutine)

	wg.Add(6)
	go a.loadAbis(&wg, errorChan)
	go a.loadManifest(&wg, errorChan)
	go a.loadMonitors(&wg, errorChan)
	go a.loadIndex(&wg, errorChan)
	go a.loadStatus(&wg, errorChan)
	go a.loadSettings(&wg, errorChan)

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

	a.loadProject(nil, nil)

	if len(errors) > 0 {
		// Handle errors, e.g., wait 1/2 second between each error message
		for _, err := range errors {
			messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
				String1: err.Error(),
			})
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		messages.EmitMessage(a.ctx, messages.Daemon, &messages.MessageMsg{
			Name:    a.FreshenController.Name,
			String1: "Freshening...",
			String2: a.FreshenController.Color,
		})
	}
	return nil
}
