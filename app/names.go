package app

import (
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) GetNames(first, pageSize int) types.SummaryName {
	first = base.Max(0, base.Min(first, len(a.names.Names)-1))
	last := base.Min(len(a.names.Names), first+pageSize)
	copy := a.names.ShallowCopy()
	copy.Names = a.names.Names[first:last]
	return copy
}

func (a *App) GetNamesCnt() int {
	return len(a.names.Names)
}

func (a *App) loadNames() error {
	nameTypes := []coreTypes.Parts{coreTypes.Regular, coreTypes.Custom, coreTypes.Prefund, coreTypes.Baddress}
	for _, t := range nameTypes {
		if namesMap, err := names.LoadNamesMap("mainnet", t, nil); err != nil {
			return err
		} else {
			for addr, name := range namesMap {
				name.Parts |= t
				if vv, ok := a.names.NamesMap[addr]; ok {
					name = vv
					name.Parts |= t
				}
				a.names.NamesMap[addr] = name
			}
		}
	}
	for _, name := range a.names.NamesMap {
		a.names.Names = append(a.names.Names, name)
	}
	sort.Slice(a.names.Names, func(i, j int) bool {
		ti := a.names.Names[i].Parts
		if ti == coreTypes.Regular {
			ti = 7
		}
		tj := a.names.Names[j].Parts
		if tj == coreTypes.Regular {
			tj = 7
		}
		if ti == tj {
			if a.names.Names[i].Tags == a.names.Names[j].Tags {
				return a.names.Names[i].Address.Hex() < a.names.Names[j].Address.Hex()
			}
			return a.names.Names[i].Tags < a.names.Names[j].Tags
		}
		return ti < tj
	})
	a.names.Summarize()
	return nil
}
