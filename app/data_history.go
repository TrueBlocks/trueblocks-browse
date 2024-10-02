package app

import (
	"fmt"
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) HistoryPage(addr string, first, pageSize int) *types.HistoryContainer {
	// logger.Info("Getting history page for", addr, "from", first, "to", first+pageSize)
	if !a.isConfigured() {
		// logger.Info("Not configured")
		return &types.HistoryContainer{}
	}

	address, ok := a.ConvertToAddress(addr)
	if !ok {
		// logger.Error("Invalid address: " + addr)
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(fmt.Errorf("Invalid address: "+addr)))
		return &types.HistoryContainer{}
	}

	// historyMutex.RLock()
	_, exists := a.project.HistoryMap.Load(address)
	// historyMutex.RUnlock()

	// logger.Info("Address okay. Exists:", exists)
	if !exists {
		messages.Send(a.ctx,
			messages.Progress,
			messages.NewProgressMsg(0, 0, address),
		)

		rCtx := a.RegisterCtx(address)
		opts := sdk.ExportOptions{
			Addrs:     []string{addr},
			RenderCtx: rCtx,
			// Articulate: true,
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
					summary.Address = address
					summary.Name = a.names.NamesMap[address].Name
					summary.Items = append(summary.Items, *tx)
					if len(summary.Items)%base.Max(pageSize, 1) == 0 {
						sort.Slice(summary.Items, func(i, j int) bool {
							if summary.Items[i].BlockNumber == summary.Items[j].BlockNumber {
								return summary.Items[i].TransactionIndex > summary.Items[j].TransactionIndex
							}
							return summary.Items[i].BlockNumber > summary.Items[j].BlockNumber
						})
						messages.Send(a.ctx,
							messages.Progress,
							messages.NewProgressMsg(int64(len(summary.Items)), int64(nItems), address),
						)
					}

					if len(summary.Items) == 0 {
						a.project.HistoryMap.Delete(address)
					} else {
						a.project.HistoryMap.Store(address, summary)
					}

				case err := <-opts.RenderCtx.ErrorChan:
					// logger.Error("Error getting history 1: " + err.Error())
					messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))

				default:
					if opts.RenderCtx.WasCanceled() {
						// logger.Warn("Load was canceled: " + address.Hex())
						return
					}
				}
			}
		}()

		_, meta, err := opts.Export() // blocks until forever loop above finishes
		if err != nil {
			// logger.Error("Error getting history 2: " + err.Error())
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))
			return &types.HistoryContainer{}
		}
		a.meta = *meta

		// logger.Info("Finished export")

		summary, _ := a.project.HistoryMap.Load(address)
		sort.Slice(summary.Items, func(i, j int) bool {
			if summary.Items[i].BlockNumber == summary.Items[j].BlockNumber {
				return summary.Items[i].TransactionIndex > summary.Items[j].TransactionIndex
			}
			return summary.Items[i].BlockNumber > summary.Items[j].BlockNumber
		})
		summary.Summarize()
		a.project.HistoryMap.Store(address, summary)

		messages.Send(a.ctx,
			messages.Completed,
			messages.NewProgressMsg(int64(a.txCount(address)), int64(a.txCount(address)), address),
		)

		// logger.Info("Loading project")
		a.loadProject(nil, nil)
	}

	if first == -1 {
		// logger.Info("Leaving...")
		return &types.HistoryContainer{}
	}

	first = base.Max(0, base.Min(first, a.txCount(address)-1))
	last := base.Min(a.txCount(address), first+pageSize)
	sum, _ := a.project.HistoryMap.Load(address)
	sum.Summarize()
	copy := sum.ShallowCopy().(*types.HistoryContainer)
	copy.Balance = a.getBalance(address)
	copy.Items = sum.Items[first:last]
	// logger.Info("Returning history page for", addr, "from", first, "to", last)
	return copy
}

func (a *App) getHistoryCnt(address base.Address) int {
	opts := sdk.ListOptions{
		Addrs:   []string{address.Hex()},
		Globals: a.globals,
	}
	appearances, meta, err := opts.ListCount()
	if err != nil {
		messages.SendError(a.ctx, err, address)
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
		if !process(&value) {
			return false
		}
		return true
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

func (a *App) closeFile(address base.Address) {
	a.project.HistoryMap.Delete(address)
}

func (a *App) txCount(address base.Address) int {
	if a.isFileOpen(address) {
		history, _ := a.project.HistoryMap.Load(address)
		return len(history.Items)
	} else {
		return 0
	}
}
