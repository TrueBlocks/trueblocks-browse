package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
)

var ctxMutex sync.Mutex

func (a *App) registerCtx(address base.Address) *output.RenderCtx {
	ctxMutex.Lock()
	defer ctxMutex.Unlock()

	rCtx := output.NewStreamingContext()
	a.renderCtxs[address] = append(a.renderCtxs[address], rCtx)
	return rCtx
}

func (a *App) unregisterCtx(address base.Address) bool {
	return a.cancelContext(address)
}

func (a *App) cancelContext(address base.Address) (removed bool) {
	ctxMutex.Lock()
	defer ctxMutex.Unlock()
	if ctxArrays, ok := a.renderCtxs[address]; ok {
		for _, ctx := range ctxArrays {
			a.emitMsg(messages.Cancelled, &messages.MessageMsg{Address: address})
			ctx.Cancel()
		}
		delete(a.renderCtxs, address)
		removed = true
	}
	return
}

func (a *App) CancelAllContexts() {
	for address := range a.renderCtxs {
		a.cancelContext(address)
	}
}
