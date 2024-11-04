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

// Freshen when the app starts and then later by the daemons to instruct the backend and
// by extension the frontend to update. We protect against updating too fast... Note
// that this routine is called as a goroutine.
func (a *App) Freshen() error {
	if !a.isConfigured() {
		return fmt.Errorf("App not configured")
	}

	if !freshenLock.CompareAndSwap(0, 1) {
		return nil // it's okay to skip a refresh if one is already in progress
	}
	defer freshenLock.CompareAndSwap(1, 0)

	freshenMutex.Lock()
	defer freshenMutex.Unlock()

	wg := sync.WaitGroup{}
	errorChan := make(chan error, 5) // Buffered channel to hold up to 5 errors (one from each goroutine)

	logger.InfoBB("")
	logger.InfoBB("--------------------- Freshen ---------------------", time.Now().Format("15:04:05"))

	// Always make sure names are loaded. We need them throughout (put any errors in the errorChan).
	_ = a.loadNames(nil, errorChan)

	// The rest of the data is independant of each other and may be loaded in parallel
	wg.Add(9)

	// app/data_project.go:
	go a.loadProjects(&wg, errorChan)

	// app/data_history.go:
	go a.loadHistory(a.GetSelected(), &wg, errorChan)

	// app/data_monitor.go:
	go a.loadMonitors(&wg, errorChan)

	// app/data_name.go:
	// go a.loadNames(&wg, errorChan)

	// app/data_abi.go:
	go a.loadAbis(&wg, errorChan)

	// app/data_index.go:
	go a.loadIndexes(&wg, errorChan)

	// app/data_manifest.go:
	go a.loadManifests(&wg, errorChan)

	// app/data_status.go:
	go a.loadStatus(&wg, errorChan)

	// app/data_session.go:
	go a.loadSessions(&wg, errorChan)

	// go a.loadConfig(&wg, errorChan)

	// go a.loadDaemons(&wg, errorChan)

	// go a.loadWizard(&wg, errorChan)

	// app/data_settings.go:
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

	if len(errors) > 0 {
		// Handle errors, e.g., wait 1/2 second between each error message
		for _, err := range errors {
			a.emitErrorMsg(err, nil)
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		a.emitMsg(messages.Daemon, &messages.MessageMsg{
			Name:    a.freshenController.Name,
			String1: "Freshen...",
			String2: a.freshenController.Color,
		})
	}
	return nil
}
