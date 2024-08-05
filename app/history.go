package app

import (
	"fmt"
	"strings"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: This should be on the App and it should be a sync.Map because it
// TODO: has the attributes described in the library's comments.
var addrToHistoryMap = map[base.Address][]types.TransactionEx{}
var m = sync.Mutex{}

func (a *App) GetHistoryPage(addr string, first, pageSize int) []types.TransactionEx {
	address, ok := a.ConvertToAddress(addr)
	if !ok {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(fmt.Errorf("Invalid address: "+addr)))
		return []types.TransactionEx{}
	}

	m.Lock()
	_, exists := addrToHistoryMap[address]
	m.Unlock()

	if !exists {
		rCtx := a.RegisterCtx(address)
		opts := sdk.ExportOptions{
			Addrs:     []string{addr},
			RenderCtx: rCtx,
			Globals: sdk.Globals{
				Cache: true,
				Ether: true,
			},
		}

		go func() {
			nItems := a.GetHistoryCnt(addr)
			for {
				select {
				case model := <-opts.RenderCtx.ModelChan:
					tx, ok := model.(*coreTypes.Transaction)
					if !ok {
						continue
					}
					txEx := types.NewTransactionEx(tx)
					if name, ok := a.namesSum.NamesMap[tx.From]; ok {
						txEx.FromName = name.Name
					}
					if name, ok := a.namesSum.NamesMap[tx.To]; ok {
						txEx.ToName = name.Name
					}
					m.Lock()
					addrToHistoryMap[address] = append(addrToHistoryMap[address], *txEx)
					if len(addrToHistoryMap[address])%pageSize == 0 {
						messages.Send(a.ctx,
							messages.Progress,
							messages.NewProgressMsg(int64(len(addrToHistoryMap[address])), nItems, address),
						)
					}
					m.Unlock()
				case err := <-opts.RenderCtx.ErrorChan:
					messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))
				default:
					if opts.RenderCtx.WasCanceled() {
						return
					}
				}
			}
		}()

		_, _, err := opts.Export()
		if err != nil {
			messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))
			return []types.TransactionEx{}
		}

		messages.Send(a.ctx,
			messages.Completed,
			messages.NewProgressMsg(int64(len(addrToHistoryMap[address])), int64(len(addrToHistoryMap[address])), address),
		)
	}

	m.Lock()
	defer m.Unlock()
	first = base.Max(0, base.Min(first, len(addrToHistoryMap[address])-1))
	last := base.Min(len(addrToHistoryMap[address]), first+pageSize)
	return addrToHistoryMap[address][first:last]
}

func (a *App) GetHistoryCnt(addr string) int64 {
	address, ok := a.ConvertToAddress(addr)
	if !ok {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(fmt.Errorf("Invalid address: "+addr)))
		return 0
	}

	opts := sdk.ListOptions{
		Addrs: []string{addr},
	}
	appearances, _, err := opts.ListCount()
	if err != nil {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err, address))
		return 0
	} else if len(appearances) == 0 {
		return 0
	}
	return appearances[0].NRecords
}

var e sync.Mutex

func (a *App) ConvertToAddress(addr string) (base.Address, bool) {
	if !strings.HasSuffix(addr, ".eth") {
		ret := base.HexToAddress(addr)
		return ret, ret != base.ZeroAddr
	}

	e.Lock()
	defer e.Unlock()
	if ensAddr, exists := a.ensMap[addr]; exists {
		return ensAddr, true
	}

	// Try to get an ENS or return the same input
	opts := sdk.NamesOptions{
		Terms: []string{addr},
	}
	if names, _, err := opts.Names(); err != nil {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err))
		return base.ZeroAddr, false
	} else {
		if len(names) > 0 {
			a.ensMap[addr] = names[0].Address
			return names[0].Address, true
		} else {
			ret := base.HexToAddress(addr)
			return ret, ret != base.ZeroAddr
		}
	}
}
