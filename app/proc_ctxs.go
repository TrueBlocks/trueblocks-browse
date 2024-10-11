package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
)

var ctxMutex sync.Mutex

func (a *App) RegisterCtx(address base.Address) *output.RenderCtx {
	ctxMutex.Lock()
	defer ctxMutex.Unlock()

	rCtx := output.NewStreamingContext()
	a.renderCtxs[address] = append(a.renderCtxs[address], rCtx)
	return rCtx
}

func (a *App) CancelContext(address base.Address) {
	ctxMutex.Lock()
	defer ctxMutex.Unlock()
	if ctxArrays, ok := a.renderCtxs[address]; ok {
		for _, ctx := range ctxArrays {
			messages.Send(a.ctx, messages.Cancelled, messages.NewCancelMsg(address))
			ctx.Cancel()
		}
		delete(a.renderCtxs, address)
	}
}

func (a *App) CancelAllContexts() {
	for address := range a.renderCtxs {
		a.CancelContext(address)
	}
}
