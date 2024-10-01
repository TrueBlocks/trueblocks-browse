package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
)

var ctxMutex sync.Mutex

func (a *App) RegisterCtx(addr base.Address) *output.RenderCtx {
	ctxMutex.Lock()
	defer ctxMutex.Unlock()

	rCtx := output.NewStreamingContext()
	a.renderCtxs[addr] = append(a.renderCtxs[addr], rCtx)
	return rCtx
}

func (a *App) CancelContexts() {
	for address, ctxArrays := range a.renderCtxs {
		for _, ctx := range ctxArrays {
			messages.Send(a.ctx,
				messages.Cancelled,
				messages.NewProgressMsg(int64(len(a.historyMap[address].Items)), int64(len(a.historyMap[address].Items)), address),
			)
			ctx.Cancel()
		}
		delete(a.renderCtxs, address)
	}
}
