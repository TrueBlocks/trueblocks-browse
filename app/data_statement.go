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

var statementsLock atomic.Uint32

func (a *App) loadStatements(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadStatements", false)()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !statementsLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer statementsLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	// EXISTING_CODE

	if !a.isConfigured() || !a.statements.NeedsUpdate() {
		return nil
	}
	updater := a.statements.Updater
	defer func() {
		a.statements.Updater = updater
	}()
	logger.InfoBY("Updating statements...")

	if items, meta, err := a.pullStatements(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		err = fmt.Errorf("no statements found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		// EXISTING_CODE
		// EXISTING_CODE
		a.meta = *meta
		a.statements = types.NewStatementContainer(a.getChain(), items, a.GetLastAddress())
		// EXISTING_CODE
		// EXISTING_CODE
		if err := a.statements.Sort(); err != nil {
			a.emitErrorMsg(err, nil)
		}
		a.emitLoadingMsg(messages.Loaded, "statements")
	}

	return nil
}

func (a *App) pullStatements() (items []types.Transaction, meta *types.Meta, err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
