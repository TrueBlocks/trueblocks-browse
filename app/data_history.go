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
		messages.EmitError(a.ctx, fmt.Errorf("Invalid address: "+addr))
		return &types.HistoryContainer{}
	}

	_, exists := a.project.HistoryMap.Load(address)
	if !exists {
		return &types.HistoryContainer{}
	}

	first = base.Max(0, base.Min(first, a.txCount(address)-1))
	last := base.Min(a.txCount(address), first+pageSize)
	history, _ := a.project.HistoryMap.Load(address)
	history.Summarize()
	copy := history.ShallowCopy().(*types.HistoryContainer)
	copy.Balance = a.getBalance(address)
	copy.Items = history.Items[first:last]
	return copy
}

func (a *App) getHistoryCnt(address base.Address) int {
	opts := sdk.ListOptions{
		Addrs:   []string{address.Hex()},
		Globals: a.globals,
	}
	appearances, meta, err := opts.ListCount()
	if err != nil {
		messages.EmitError(a.ctx, err, address)
		return 0
	} else if len(appearances) == 0 {
		return 0
	} else {
		a.meta = *meta
		return int(appearances[0].NRecords)
	}
}

func (a *App) forEveryTx(address base.Address, process func(coreTypes.Transaction) bool) bool {
	historyContainer, _ := a.project.HistoryMap.Load(address)
	for _, item := range historyContainer.Items {
		if !process(item) {
			return false
		}
	}
	return true
}

func (a *App) forEveryHistory(process func(*types.HistoryContainer) bool) bool {
	a.project.HistoryMap.Range(func(key base.Address, value types.HistoryContainer) bool {
		return process(&value)
	})
	return true
}

func (a *App) isFileOpen(address base.Address) bool {
	_, isOpen := a.project.HistoryMap.Load(address)
	return isOpen
}

func (a *App) openFileCnt() int {
	count := 0
	a.project.HistoryMap.Range(func(key base.Address, value types.HistoryContainer) bool {
		count++
		return true
	})
	return count
}

func (a *App) txCount(address base.Address) int {
	if a.isFileOpen(address) {
		history, _ := a.project.HistoryMap.Load(address)
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

	history, exists := a.project.HistoryMap.Load(address)
	if exists {
		if !history.NeedsUpdate(a.nameChange()) {
			return nil
		}
	}

	logger.Info("Loading history for address: ", address.Hex())
	if err := a.thing(address, 15); err != nil {
		messages.EmitError(a.ctx, err, address)
		return err
	}
	a.loadProject(nil, nil)

	return nil
}

func (a *App) thing(address base.Address, freq int) error {
	rCtx := a.RegisterCtx(address)
	opts := sdk.ExportOptions{
		Addrs:     []string{address.Hex()},
		RenderCtx: rCtx,
		Globals: sdk.Globals{
			Cache: true,
			Ether: true,
			Chain: a.globals.Chain,
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
				summary, _ := a.project.HistoryMap.Load(address)
				summary.NTotal = nItems
				summary.Address = address
				summary.Name = a.names.NamesMap[address].Name
				summary.Items = append(summary.Items, *tx)
				if len(summary.Items)%freq == 0 {
					sort.Slice(summary.Items, func(i, j int) bool {
						if summary.Items[i].BlockNumber == summary.Items[j].BlockNumber {
							return summary.Items[i].TransactionIndex > summary.Items[j].TransactionIndex
						}
						return summary.Items[i].BlockNumber > summary.Items[j].BlockNumber
					})
					messages.EmitProgress(a.ctx, address, len(summary.Items), nItems)
				}

				if len(summary.Items) == 0 {
					a.project.HistoryMap.Delete(address)
				} else {
					a.project.HistoryMap.Store(address, summary)
				}

			case err := <-opts.RenderCtx.ErrorChan:
				messages.EmitError(a.ctx, err, address)

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

	history, _ := a.project.HistoryMap.Load(address)
	sort.Slice(history.Items, func(i, j int) bool {
		if history.Items[i].BlockNumber == history.Items[j].BlockNumber {
			return history.Items[i].TransactionIndex > history.Items[j].TransactionIndex
		}
		return history.Items[i].BlockNumber > history.Items[j].BlockNumber
	})
	history.Summarize()
	a.project.HistoryMap.Store(address, history)
	messages.EmitCompleted(a.ctx, address, a.txCount(address))
	return nil
}
