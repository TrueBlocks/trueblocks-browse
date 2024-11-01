package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) ConvertToAddress(addr string) (base.Address, bool) {
	if !a.IsConfigured() {
		return base.ZeroAddr, false
	}

	if !strings.HasSuffix(addr, ".eth") {
		ret := base.HexToAddress(addr)
		return ret, ret != base.ZeroAddr
	}

	ensAddr, exists := a.EnsCache.Load(addr)
	if exists {
		return ensAddr.(base.Address), true
	}

	// Try to get an ENS or return the same input
	opts := sdk.NamesOptions{
		Terms:   []string{addr},
		Globals: a.toGlobals(),
	}
	if names, meta, err := opts.Names(); err != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
		})
		return base.ZeroAddr, false
	} else {
		a.meta = *meta
		if len(names) > 0 {
			a.EnsCache.Store(addr, names[0].Address)
			return names[0].Address, true
		} else {
			ret := base.HexToAddress(addr)
			return ret, ret != base.ZeroAddr
		}
	}
}

func (a *App) AddrToName(address base.Address) string {
	if name, exists := a.names.NamesCache[address]; exists {
		return name.Name
	}
	return ""
}
