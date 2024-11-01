package app

import (
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

func (a *App) ConvertToAddress(addr string) (base.Address, bool) {
	if !a.isConfigured() {
		return base.ZeroAddr, false
	}

	if !strings.HasSuffix(addr, ".eth") {
		ret := base.HexToAddress(addr)
		return ret, ret != base.ZeroAddr
	}

	ensAddr, exists := a.ensCache.Load(addr)
	if exists {
		return ensAddr.(base.Address), true
	}

	// Try to get an ENS or return the same input
	opts := sdk.NamesOptions{
		Terms:   []string{addr},
		Globals: a.toGlobals(),
	}
	if names, meta, err := opts.Names(); err != nil {
		a.emitErrorMsg(err, nil)
		return base.ZeroAddr, false
	} else {
		a.meta = *meta
		if len(names) > 0 {
			a.ensCache.Store(addr, names[0].Address)
			return names[0].Address, true
		} else {
			ret := base.HexToAddress(addr)
			return ret, ret != base.ZeroAddr
		}
	}
}

func (a *App) AddrToName(address base.Address) string {
	if name, exists := a.namesMap[address]; exists {
		return name.Name
	}
	return ""
}
