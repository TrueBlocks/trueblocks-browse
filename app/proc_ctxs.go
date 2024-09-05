package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
)

var r sync.Mutex

func (a *App) RegisterCtx(addr base.Address) *output.RenderCtx {
	r.Lock()
	defer r.Unlock()

	rCtx := output.NewStreamingContext()
	a.renderCtxs[addr] = append(a.renderCtxs[addr], rCtx)
	return rCtx
}

func (a *App) Cancel(addr base.Address) (int, bool) {
	if len(a.renderCtxs) == 0 {
		return 0, false
	}
	if a.renderCtxs[addr] == nil {
		return 0, true
	}
	n := len(a.renderCtxs[addr])
	for i := 0; i < len(a.renderCtxs[addr]); i++ {
		a.renderCtxs[addr][i].Cancel()
	}
	a.renderCtxs[addr] = nil
	return n, true
}
