// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var historyLock atomic.Uint32

func (a *App) loadHistory(address base.Address, wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadHistory")()
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !historyLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer historyLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	_ = a.forceHistory()
	if address.IsZero() {
		return nil
	}
	// EXISTING_CODE

	history, exists := a.historyCache.Load(address)
	if exists && len(history.Items) > 0 {
		// if !history.NeedsUpdate(a.forceHistory()) {
		return nil // we only update with a Reload
		// }
	}

	if err := a.thing(address, 150, errorChan); err != nil {
		a.emitAddressErrorMsg(err, address)
		return err
	}

	return nil
}

func (a *App) forceHistory() (force bool) {
	// EXISTING_CODE
	force = a.forceName()
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (a *App) thing(address base.Address, freq int, errorChan chan error) error {
	txCnt := a.txCount(address)
	a.emitProgressMsg(messages.Started, address, 0, 0)
	defer func() {
		a.emitProgressMsg(messages.Completed, address, txCnt, txCnt)
	}()
	_ = errorChan // delint
	rCtx := a.registerCtx(address)
	defer a.unregisterCtx(address)

	opts := sdk.ExportOptions{
		Addrs:     []string{address.Hex()},
		RenderCtx: rCtx,
		Globals: sdk.Globals{
			Cache: true,
			Ether: true,
			Chain: a.getChain(),
		},
	}

	expectedCnt := a.getHistoryCnt(address)

	// we always have a currently opened history which contains the entire history so far...
	history, _ := a.historyCache.Load(address)
	history.NTotal = uint64(expectedCnt)
	history.Address = address
	history.Name = a.namesMap[address].Name
	go func() {
		for {
			select {
			case model := <-opts.RenderCtx.ModelChan:
				tx, ok := model.(*coreTypes.Transaction)
				if !ok {
					continue
				}
				history.Items = append(history.Items, *tx)
				// ...periodically, we store it back into the map...
				if len(history.Items)%freq == 0 {
					// sort it first...
					sort.Slice(history.Items, func(i, j int) bool {
						if history.Items[i].BlockNumber == history.Items[j].BlockNumber {
							return history.Items[i].TransactionIndex > history.Items[j].TransactionIndex
						}
						return history.Items[i].BlockNumber > history.Items[j].BlockNumber
					})
					// put it back into the history cache so other processes can use it...
					a.historyCache.Store(address, history)
				}
				// ...let the front end know and keep going...(note we use the same history)
				a.emitProgressMsg(messages.Progress, address, len(history.Items), int(expectedCnt))

			case err := <-opts.RenderCtx.ErrorChan:
				a.emitAddressErrorMsg(err, address)

			default:
				if opts.RenderCtx.WasCanceled() {
					return
				}
			}
		}
	}()

	_, meta, err := opts.Export() // blocks until forever loop above finishes
	if err != nil {
		a.emitProgressMsg(messages.Canceled, address, txCnt, txCnt)
		return err
	}

	txCnt = a.txCount(address)
	a.meta = *meta
	sort.Slice(history.Items, func(i, j int) bool {
		if history.Items[i].BlockNumber == history.Items[j].BlockNumber {
			return history.Items[i].TransactionIndex > history.Items[j].TransactionIndex
		}
		return history.Items[i].BlockNumber > history.Items[j].BlockNumber
	})
	a.historyCache.Store(address, history)

	return nil
}

func (a *App) Reload() {
	defer a.trackPerformance("Reload")()

	switch a.session.LastRoute {
	case "/names":
		logger.InfoG("Reloading names...")
		a.names.LastUpdate = 0
		if err := a.loadNames(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	default:
		logger.InfoG("Reloading default (history)...")
		history, _ := a.historyCache.Load(a.GetSelected())
		history.LastUpdate = 0
		a.historyCache.Store(a.GetSelected(), history)
		a.goToAddress(history.Address)
	}
}

func (a *App) getHistoryCnt(address base.Address) int64 {
	fn := coreMonitor.PathToMonitorFile(a.getChain(), address)
	return (file.FileSize(fn) / index.AppRecordWidth) - 1 // header
}

func (a *App) txCount(address base.Address) int {
	if history, exists := a.historyCache.Load(address); exists {
		return len(history.Items)
	} else {
		return 0
	}
}

func (a *App) goToAddress(address base.Address) {
	defer a.trackPerformance("goToAddress")()

	if address == base.ZeroAddr {
		return
	}

	a.cancelContext(address)
	a.SetRoute("/history", address.Hex())
	history, exists := a.historyCache.Load(address)
	if exists {
		history.Items = make([]coreTypes.Transaction, 0, len(history.Items))
		a.historyCache.Store(address, history)
	}
	go a.loadHistory(address, nil, nil)
	a.emitNavigateMsg(a.GetRoute())
	a.emitMsg(messages.Refresh, &messages.MessageMsg{
		Num1: 3, // 3 means refresh
	})
}

func (a *App) LoadAddress(addrOrEns string) {
	if address, ok := a.ensToAddress(addrOrEns); ok {
		a.goToAddress(address)
	} else {
		a.emitErrorMsg(ErrInvalidAddress, nil)
	}
}

// EXISTING_CODE
