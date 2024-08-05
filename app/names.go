package app

import (
	"sort"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) GetNamesPage(first, pageSize int) []coreTypes.Name {
	if len(a.names) == 0 {
		return a.names
	}

	first = base.Max(0, base.Min(first, len(a.names)-1))
	last := base.Min(len(a.names), first+pageSize)
	return a.names[first:last]
}

func (a *App) GetNamesCnt() int {
	return len(a.names)
}

func (a *App) loadNames() error {
	nameTypes := []coreTypes.Parts{coreTypes.Regular, coreTypes.Custom, coreTypes.Prefund, coreTypes.Baddress}
	for _, t := range nameTypes {
		if namesMap, err := names.LoadNamesMap("mainnet", t, nil); err != nil {
			return err
		} else {
			for addr, name := range namesMap {
				name.Parts |= t
				if vv, ok := a.namesMap[addr]; ok {
					name = vv
					name.Parts |= t
				}
				a.namesMap[addr] = name
			}
		}
	}
	for _, name := range a.namesMap {
		a.names = append(a.names, name)
	}
	sort.Slice(a.names, func(i, j int) bool {
		ti := a.names[i].Parts
		if ti == coreTypes.Regular {
			ti = 7
		}
		tj := a.names[j].Parts
		if tj == coreTypes.Regular {
			tj = 7
		}
		if ti == tj {
			if a.names[i].Tags == a.names[j].Tags {
				return a.names[i].Address.Hex() < a.names[j].Address.Hex()
			}
			return a.names[i].Tags < a.names[j].Tags
		}
		return ti < tj
	})
	return nil
}
