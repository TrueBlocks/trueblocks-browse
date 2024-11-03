// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) HistoryPage(addr string, first, pageSize int) *types.HistoryContainer {
	address, ok := a.ConvertToAddress(addr)
	if !ok {
		err := fmt.Errorf("Invalid address: " + addr)
		a.emitErrorMsg(err, nil)
		return &types.HistoryContainer{}
	}

	_, exists := a.historyCache.Load(address)
	if !exists {
		return &types.HistoryContainer{}
	}

	first = base.Max(0, base.Min(first, a.txCount(address)-1))
	last := base.Min(a.txCount(address), first+pageSize)
	history, _ := a.historyCache.Load(address)
	history.Summarize()
	copy := history.ShallowCopy().(*types.HistoryContainer)
	copy.Balance = a.getBalance(address)
	copy.Items = history.Items[first:last]
	return copy
}

func (a *App) getHistoryCnt(address base.Address) uint64 {
	opts := sdk.ListOptions{
		Addrs:   []string{address.Hex()},
		Globals: a.toGlobals(),
	}
	if appearances, meta, err := opts.ListCount(); err != nil {
		a.emitAddressErrorMsg(err, address)
		return 0
	} else if len(appearances) == 0 {
		return 0
	} else {
		a.meta = *meta
		return uint64(appearances[0].NRecords)
	}
}

func (a *App) isFileOpen(address base.Address) bool {
	_, isOpen := a.historyCache.Load(address)
	return isOpen
}

func (a *App) txCount(address base.Address) int {
	if a.isFileOpen(address) {
		history, _ := a.historyCache.Load(address)
		return len(history.Items)
	} else {
		return 0
	}
}

// var historyLock atomic.Uint32

func (a *App) loadHistory(address base.Address, wg *sync.WaitGroup, errorChan chan error) error {
	if wg != nil {
		defer wg.Done()
	}

	if address.IsZero() {
		return nil
	}

	// if !historyLock.CompareAndSwap(0, 1) {
	// 	return nil
	// }
	// defer historyLock.CompareAndSwap(1, 0)

	history, exists := a.historyCache.Load(address)
	if exists {
		if !history.NeedsUpdate(a.forceHistory()) {
			return nil
		}
	}

	_ = errorChan // delint
	logger.Info("Loading history for address: ", address.Hex())
	if err := a.thing(address, 15); err != nil {
		a.emitAddressErrorMsg(err, address)
		return err
	}
	a.loadProjects(nil, nil)

	return nil
}

func (a *App) thing(address base.Address, freq int) error {
	rCtx := a.registerCtx(address)
	defer a.unregisterCtx(address)

	opts := sdk.ExportOptions{
		Addrs:     []string{address.Hex()},
		RenderCtx: rCtx,
		Globals: sdk.Globals{
			Cache: true,
			Ether: true,
			Chain: a.Chain,
		},
	}

	go func() {
		nItems := a.getHistoryCnt(address)
		for {
			select {
			case model := <-opts.RenderCtx.ModelChan:
				tx, ok := model.(*coreTypes.Transaction)
				if !ok {
					continue
				}
				summary, _ := a.historyCache.Load(address)
				summary.NTotal = nItems
				summary.Address = address
				summary.Name = a.namesMap[address].Name
				summary.Items = append(summary.Items, *tx)
				if len(summary.Items)%(freq*3) == 0 {
					sort.Slice(summary.Items, func(i, j int) bool {
						if summary.Items[i].BlockNumber == summary.Items[j].BlockNumber {
							return summary.Items[i].TransactionIndex > summary.Items[j].TransactionIndex
						}
						return summary.Items[i].BlockNumber > summary.Items[j].BlockNumber
					})
					a.emitProgressMsg(messages.Progress, address, len(summary.Items), int(nItems))
				}

				if len(summary.Items) == 0 {
					a.historyCache.Delete(address)
				} else {
					a.historyCache.Store(address, summary)
				}

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
		return err
	}
	a.meta = *meta

	history, _ := a.historyCache.Load(address)
	sort.Slice(history.Items, func(i, j int) bool {
		if history.Items[i].BlockNumber == history.Items[j].BlockNumber {
			return history.Items[i].TransactionIndex > history.Items[j].TransactionIndex
		}
		return history.Items[i].BlockNumber > history.Items[j].BlockNumber
	})
	history.Summarize()
	a.historyCache.Store(address, history)
	a.emitProgressMsg(messages.Completed, address, a.txCount(address), a.txCount(address))
	return nil
}

func (a *App) forceHistory() (force bool) {
	// EXISTING_CODE
	force = a.forceName()
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (a *App) Reload() {
	switch a.session.LastRoute {
	case "/names":
		logger.InfoC("Reloading names")
		a.names.LastUpdate = time.Time{}
		if err := a.loadNames(nil, nil); err != nil {
			a.emitErrorMsg(err, nil)
		}
	}
}

func (a *App) GoToAddress(address base.Address) {
	logger.InfoW("--------------------------- enter -------------------------------------------")
	logger.InfoW("GoToAddress: ", address.Hex())
	if address == base.ZeroAddr {
		logger.InfoW("--------------------------- zeroAddr exit -------------------------------------------")
		return
	}

	a.SetRoute("/history", address.Hex())

	a.cancelContext(address)
	a.historyCache.Delete(address)
	a.loadHistory(address, nil, nil)

	a.emitNavigateMsg(a.GetRoute())
	a.emitInfoMsg("viewing address", address.Hex())
	logger.InfoW("--------------------------- exit -------------------------------------------")
}

// EXISTING_CODE
