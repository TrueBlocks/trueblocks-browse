package app

import (
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
		messages.SendMessage(a.ctx, base.ZeroAddr, messages.Error, "Invalid address")
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
					txEx := types.NewTransactionEx(a.namesMap, tx)
					m.Lock()
					addrToHistoryMap[address] = append(addrToHistoryMap[address], *txEx)
					if len(addrToHistoryMap[address])%pageSize == 0 {
						messages.SendMessage(a.ctx, address, messages.Progress, &messages.ProgressMsg{
							Address: address,
							Have:    int64(len(addrToHistoryMap[address])),
							Want:    nItems,
						})
					}
					m.Unlock()
				case err := <-opts.RenderCtx.ErrorChan:
					messages.SendMessage(a.ctx, address, messages.Error, err.Error())
				default:
					if opts.RenderCtx.WasCanceled() {
						return
					}
				}
			}
		}()

		_, _, err := opts.Export()
		if err != nil {
			messages.SendMessage(a.ctx, address, messages.Error, err.Error())
			return []types.TransactionEx{}
		}

		messages.SendMessage(a.ctx, address, messages.Completed, "")
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
		messages.SendMessage(a.ctx, base.ZeroAddr, messages.Error, "Invalid address")
		return 0
	}

	opts := sdk.ListOptions{
		Addrs: []string{addr},
	}
	monitors, _, err := opts.ListCount()
	if err != nil {
		messages.SendMessage(a.ctx, address, messages.Error, err.Error())
		return 0
	} else if len(monitors) == 0 {
		return 0
	}
	return monitors[0].NRecords
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
		messages.SendMessage(a.ctx, base.ZeroAddr, messages.Error, err.Error())
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
