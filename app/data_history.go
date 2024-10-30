// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

import (
	"fmt"
	"sort"
	"sync"

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
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
		})
		return &types.HistoryContainer{}
	}

	_, exists := a.projects.HistoryMap.Load(address)
	if !exists {
		return &types.HistoryContainer{}
	}

	first = base.Max(0, base.Min(first, a.txCount(address)-1))
	last := base.Min(a.txCount(address), first+pageSize)
	history, _ := a.projects.HistoryMap.Load(address)
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
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
			Address: address,
		})
		return 0
	} else if len(appearances) == 0 {
		return 0
	} else {
		a.meta = *meta
		return uint64(appearances[0].NRecords)
	}
}

func (a *App) forEveryHistory(process func(*types.HistoryContainer) bool) bool {
	a.projects.HistoryMap.Range(func(key base.Address, value types.HistoryContainer) bool {
		return process(&value)
	})
	return true
}

func (a *App) isFileOpen(address base.Address) bool {
	_, isOpen := a.projects.HistoryMap.Load(address)
	return isOpen
}

func (a *App) openFileCnt() int {
	count := 0
	a.projects.HistoryMap.Range(func(key base.Address, value types.HistoryContainer) bool {
		count++
		return true
	})
	return count
}

func (a *App) txCount(address base.Address) int {
	if a.isFileOpen(address) {
		history, _ := a.projects.HistoryMap.Load(address)
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

	history, exists := a.projects.HistoryMap.Load(address)
	if exists {
		if !history.NeedsUpdate(a.forceHistory()) {
			return nil
		}
	}

	_ = errorChan // delint
	logger.Info("Loading history for address: ", address.Hex())
	if err := a.thing(address, 15); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
			Address: address,
		})
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
				summary, _ := a.projects.HistoryMap.Load(address)
				summary.NTotal = nItems
				summary.Address = address
				summary.Name = a.names.NamesMap[address].Name
				summary.Items = append(summary.Items, *tx)
				if len(summary.Items)%(freq*3) == 0 {
					sort.Slice(summary.Items, func(i, j int) bool {
						if summary.Items[i].BlockNumber == summary.Items[j].BlockNumber {
							return summary.Items[i].TransactionIndex > summary.Items[j].TransactionIndex
						}
						return summary.Items[i].BlockNumber > summary.Items[j].BlockNumber
					})
					messages.EmitMessage(a.ctx, messages.Progress, &messages.MessageMsg{
						Address: address,
						Num1:    len(summary.Items),
						Num2:    int(nItems),
					})
				}

				if len(summary.Items) == 0 {
					a.projects.HistoryMap.Delete(address)
				} else {
					a.projects.HistoryMap.Store(address, summary)
				}

			case err := <-opts.RenderCtx.ErrorChan:
				messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
					String1: err.Error(),
					Address: address,
				})

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

	history, _ := a.projects.HistoryMap.Load(address)
	sort.Slice(history.Items, func(i, j int) bool {
		if history.Items[i].BlockNumber == history.Items[j].BlockNumber {
			return history.Items[i].TransactionIndex > history.Items[j].TransactionIndex
		}
		return history.Items[i].BlockNumber > history.Items[j].BlockNumber
	})
	history.Summarize()
	a.projects.HistoryMap.Store(address, history)
	messages.EmitMessage(a.ctx, messages.Completed, &messages.MessageMsg{
		Address: address,
		Num1:    a.txCount(address),
		Num2:    a.txCount(address),
	})
	return nil
}

func (a *App) forceHistory() (force bool) {
	// EXISTING_CODE
	force = a.forceName()
	// EXISTING_CODE
	return
}
