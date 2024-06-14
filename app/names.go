package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/sdk"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) GetNames(first, pageSize int) []string {
	var ret []string
	if len(a.names) == 0 {
		return ret
	}
	first = base.Max(0, base.Min(first, len(a.names)-1))
	last := base.Min(len(a.names), first+pageSize)
	n := a.names[first:last]
	for _, name := range n {
		ret = append(ret, fmt.Sprintf("%s %s %s", name.Address.Hex(), name.Tags, name.Name))
	}
	return ret
}

func (a *App) MaxNames() int {
	return len(a.names)
}

func (a *App) loadNames() ([]types.Name, error) {
	opts := sdk.NamesOptions{
		Regular: true,
		// Custom:  true,
		Globals: sdk.Globals{
			Chain: "mainnet",
		},
	}
	names, _, err := opts.Names()
	for i := range names {
		names[i].Name = fmt.Sprintf("%d: %s", i, names[i].Name)
	}
	return names, err
}
