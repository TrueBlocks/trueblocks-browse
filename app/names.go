package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/sdk"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) GetNames(first, pageSize int) []string {
	last := first + pageSize
	if len(a.names) < last {
		return []string{"No names loaded"}
	}
	n := a.names[first:last]
	var ret []string
	for _, name := range n {
		ret = append(ret, fmt.Sprintf("%s: %s", name.Address.Hex(), name.Name))
	}
	return ret
}

func (a *App) loadNames() ([]types.Name, error) {
	opts := sdk.NamesOptions{
		Custom:  true,
		Prefund: true,
		Globals: sdk.Globals{
			Chain: "mainnet",
		},
	}
	names, _, err := opts.Names()
	return names, err
}
