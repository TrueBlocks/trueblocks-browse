package app

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

var freshenLock atomic.Uint32

// Freshen when the app starts and then later by the daemons to instruct the backend and
// by extension the frontend to update. We protect against updating too fast... Note
// that this routine is called as a goroutine.
func (a *App) Freshen() error {
	if !freshenLock.CompareAndSwap(0, 1) {
		return nil // it's okay to skip a refresh if one is already in progress
	}
	defer freshenLock.CompareAndSwap(1, 0)

	if !a.isConfigured() {
		return fmt.Errorf("App not configured")
	}

	// If the index is not initialized yet, we can't do anything
	if configErr := index.IsInitialized(a.getChain(), config.ExpectedVersion()); configErr != nil {
		if configErr == index.ErrNotInitialized {
			return nil
		} else {
			return configErr
		}
	}

	logger.InfoBB("")
	logger.InfoBB("--------------------- Freshen in ---------------------", time.Now().Format("15:04:05"))
	defer logger.InfoBB("--------------------- Freshen out ---------------------", time.Now().Format("15:04:05"))

	nRoutines := 9
	wg := sync.WaitGroup{}
	errorChan := make(chan error, nRoutines*2) // Buffered channel to hold up to 5 errors (one from each goroutine)

	// Always make sure names are loaded. We need them throughout (put any errors in the errorChan).
	_ = a.loadNames(nil, errorChan)

	// The rest of the data is independant of each other and may be loaded in parallel
	wg.Add(nRoutines)
	go a.loadProject(&wg, errorChan)
	// go a.loadHistory(&wg, errorChan)
	go a.loadMonitors(&wg, errorChan)
	go a.loadAbis(&wg, errorChan)
	go a.loadIndexes(&wg, errorChan)
	go a.loadManifests(&wg, errorChan)
	// go a.loadStatus(&wg, errorChan)
	go a.loadSession(&wg, errorChan)
	go a.loadConfig(&wg, errorChan)
	go a.loadDaemons(&wg, errorChan)
	go a.loadWizard(&wg, errorChan)

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

	go a.loadStatus(&wg, nil)

	if len(errors) > 0 {
		for _, err := range errors {
			a.emitErrorMsg(err, nil)
		}
	}

	a.emitMsg(messages.Refresh, &messages.MessageMsg{
		Name:    a.daemons.FreshenController.Name,
		String1: "Refresh...",
		String2: a.daemons.FreshenController.Color,
		Num1:    1, // 1 means daemon if we need it
		Bool:    len(errors) == 0,
	})

	return nil
}

/*
//------------------------------------------------------------
- Error Channel Placement: Define and initialize the error channel in NewApp and store it as a field in the App struct to make it accessible throughout the app's lifetime.
- Single Error Channel: Use a single error channel for centralized error handling, making it easier to log errors and optionally send them to the frontend.
- Background Error Handler: Start a background goroutine in NewApp to read from errorChan, log errors, and optionally forward them to the frontend.
- Non-blocking Execution: Avoid overlapping executions by using an atomic flag for re-entrancy in spawnLoadRoutines, allowing each run to complete before starting a new one.

package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type App struct {
	errorChan  chan error
	inProgress int32
}

func NewApp() *App {
	app := &App{
		errorChan: make(chan error, 5),
	}
	go app.handleErrors() // Start error handler
	return app
}

func (a *App) handleErrors() {
	for err := range a.errorChan {
		if err != nil {
			log.Println("Error:", err)
			// Optional: send error to frontend if needed
			// runtime.EventsEmit(context.Background(), "error", err.Error())
		}
	}
}

func (a *App) spawnLoadRoutines() {
	if !atomic.CompareAndSwapInt32(&a.inProgress, 0, 1) {
		return
	}
	defer atomic.StoreInt32(&a.inProgress, 0)

	go func() { a.errorChan <- a.loadExampleTask("Task 1") }()
	go func() { a.errorChan <- a.loadExampleTask("Task 2") }()
	go func() { a.errorChan <- a.loadExampleTask("Task 3") }()
}

func (a *App) lo adExampleTask(name string) error {
	time.Sleep(1 * time.Second) // Simulate task
	return fmt.Errorf("%s encountered an error", name)
}

func main() {
	app := NewApp()
	for {
		app.spawnLoadRoutines()
		time.Sleep(2 * time.Second)
	}
}

Key Points
-----------
- Error Channel in App: Initialized in NewApp, used across the app.
- Background Error Handler: handleErrors reads from errorChan and logs errors.
- Re-entrancy Control: spawnLoadRoutines uses inProgress flag to prevent overlap.
//------------------------------------------------------------
*/
