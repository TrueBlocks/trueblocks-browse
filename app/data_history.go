package app

import (
	"fmt"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

var historyMutex sync.RWMutex

func (a *App) Reload(addr base.Address) {
	a.CancleContexts()
	historyMutex.Lock()
	delete(a.historyMap, addr)
	historyMutex.Unlock()
	a.HistoryPage(addr.String(), 0, 15)
	a.removeAddress(addr)
	a.Refresh(false)
}

func (a *App) CancleContexts() {
	for address, ctxArrays := range a.renderCtxs {
		for _, ctx := range ctxArrays {
			messages.Send(a.ctx,
				messages.Cancelled,
				messages.NewProgressMsg(int64(len(a.historyMap[address].Items)), int64(len(a.historyMap[address].Items)), address),
			)
			(*ctx).Cancel()
		}
		delete(a.renderCtxs, address)
	}
}

func (a *App) HistoryPage(addr string, first, pageSize int) types.HistoryContainer {
	if !a.isConfigured() {
		return types.HistoryContainer{}
	}

	address, ok := a.ConvertToAddress(addr)
	if !ok {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(fmt.Errorf("Invalid address: "+addr)))
		return types.HistoryContainer{}
	}

	historyMutex.RLock()
	_, exists := a.historyMap[address]
	historyMutex.RUnlock()

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
			nItems := a.getHistoryCnt(addr)
			for {
				select {
				case model := <-opts.RenderCtx.ModelChan:
					tx, ok := model.(*coreTypes.Transaction)
					if !ok {
						continue
					}
					historyMutex.RLock()
					summary := a.historyMap[address]
					historyMutex.RUnlock()

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

					historyMutex.Lock()
					a.historyMap[address] = summary
					historyMutex.Unlock()

				case err := <-opts.RenderCtx.ErrorChan:
					messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))

				default:
					if opts.RenderCtx.WasCanceled() {
						return
					}
				}
			}
		}()

		_, meta, err := opts.Export() // blocks until forever loop above finishes
		if err != nil {
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))
			return types.HistoryContainer{}
		}
		a.meta = *meta

		historyMutex.Lock()
		sort.Slice(a.historyMap[address].Items, func(i, j int) bool {
			if a.historyMap[address].Items[i].BlockNumber == a.historyMap[address].Items[j].BlockNumber {
				return a.historyMap[address].Items[i].TransactionIndex > a.historyMap[address].Items[j].TransactionIndex
			}
			return a.historyMap[address].Items[i].BlockNumber > a.historyMap[address].Items[j].BlockNumber
		})
		historyMutex.Unlock()

		messages.Send(a.ctx,
			messages.Completed,
			messages.NewProgressMsg(int64(len(a.historyMap[address].Items)), int64(len(a.historyMap[address].Items)), address),
		)

		a.loadPortfolio(nil, nil)
	}

	historyMutex.RLock()
	defer historyMutex.RUnlock()

	first = base.Max(0, base.Min(first, len(a.historyMap[address].Items)-1))
	last := base.Min(len(a.historyMap[address].Items), first+pageSize)
	sum := a.historyMap[address]
	sum.Summarize()
	copy := sum.ShallowCopy()
	copy.Balance = a.getBalance(address)
	copy.Items = a.historyMap[address].Items[first:last]
	return copy
}

func (a *App) getHistoryCnt(addr string) int {
	address, ok := a.ConvertToAddress(addr)
	if !ok {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(fmt.Errorf("Invalid address: "+addr)))
		return 0
	}

	historyMutex.RLock()
	l := len(a.historyMap[address].Items)
	historyMutex.RUnlock()
	if l > 0 {
		return l
	}

	opts := sdk.ListOptions{
		Addrs:   []string{addr},
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

func (a *App) removeAddress(addr base.Address) {
	for i, item := range a.portfolio.Items {
		if item.Address == addr {
			a.portfolio.Items = append(a.portfolio.Items[:i], a.portfolio.Items[i+1:]...)
			// a.portfolio.MyCount--
			break
		}
	}
	for i, item := range a.monitors.Items {
		if item.Address == addr {
			a.monitors.Items = append(a.monitors.Items[:i], a.monitors.Items[i+1:]...)
			// a.monitors.NItems--
			break
		}
	}
}
