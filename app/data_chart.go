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
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// EXISTING_CODE

var chartsLock atomic.Uint32

func (a *App) loadCharts(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadCharts", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !chartsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer chartsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.charts.NeedsUpdate() {
		return nil
	}
	updater := a.charts.Updater
	defer func() {
		a.charts.Updater = updater
	}()
	logger.InfoBY("Updating charts...")

	if items, meta, err := a.pullCharts(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no charts found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.charts = types.NewChartContainer(a.getChain(), items, a.getLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.charts.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "charts")
	}

	return nil
}

func (a *App) pullCharts() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
