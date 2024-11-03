package app

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

var freshenLock atomic.Uint32
var freshenMutex sync.Mutex

// Refresh when the app starts and then later by the daemons to instruct the backend and
// by extension the frontend to update. We protect against updating too fast... Note
// that this routine is called as a goroutine.
func (a *App) Refresh() error {
	if !a.isConfigured() {
		return fmt.Errorf("App not configured")
	}

	if !freshenLock.CompareAndSwap(0, 1) {
		return nil // it's okay to skip a refresh if one is already in progress
	}
	defer freshenLock.CompareAndSwap(1, 0)

	freshenMutex.Lock()
	defer freshenMutex.Unlock()

	if !a.scraperController.IsRunning() {
		logger.InfoG("Freshening...")
	}

	// We always load names first since we need them everywhere
	err := a.loadNames(nil, nil)
	if err != nil {
		a.emitErrorMsg(err, nil)
	}

	// And then update everything else in the fullness of time
	wg := sync.WaitGroup{}
	errorChan := make(chan error, 5) // Buffered channel to hold up to 5 errors (one from each goroutine)

	wg.Add(8)
	go a.loadProjects(&wg, errorChan)
	go a.loadMonitors(&wg, errorChan)
	go a.loadSessions(&wg, errorChan)
	go a.loadSettings(&wg, errorChan)
	go a.loadStatus(&wg, errorChan)
	go a.loadAbis(&wg, errorChan)
	go a.loadManifests(&wg, errorChan)
	go a.loadIndexes(&wg, errorChan)

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

	if len(errors) > 0 {
		// Handle errors, e.g., wait 1/2 second between each error message
		for _, err := range errors {
			a.emitErrorMsg(err, nil)
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		a.emitMsg(messages.Daemon, &messages.MessageMsg{
			Name:    a.freshenController.Name,
			String1: "Freshening...",
			String2: a.freshenController.Color,
		})
	}
	return nil
}
