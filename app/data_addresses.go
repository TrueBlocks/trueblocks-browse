package app

import (
	"strings"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

var e sync.Mutex

func (a *App) ConvertToAddress(addr string) (base.Address, bool) {
	if !a.isConfigured() {
		return base.ZeroAddr, false
	}

	if !strings.HasSuffix(addr, ".eth") {
		ret := base.HexToAddress(addr)
		return ret, ret != base.ZeroAddr
	}

	e.Lock()
	ensAddr, exists := a.ensMap[addr]
	e.Unlock()

	if exists {
		return ensAddr, true
	}

	// Try to get an ENS or return the same input
	opts := sdk.NamesOptions{
		Terms: []string{addr},
	}
	if names, meta, err := opts.Names(); err != nil {
		messages.Send(a.ctx, messages.Error, messages.NewErrorMsg(err))
		return base.ZeroAddr, false
	} else {
		a.meta = *meta
		if len(names) > 0 {
			e.Lock()
			defer e.Unlock()
			a.ensMap[addr] = names[0].Address
			return names[0].Address, true
		} else {
			ret := base.HexToAddress(addr)
			return ret, ret != base.ZeroAddr
		}
	}
}

func (a *App) AddrToName(addr base.Address) string {
	if name, exists := a.names.NamesMap[addr]; exists {
		return name.Name
	}
	return ""
}
