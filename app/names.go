package app

import (
	"sort"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) GetNamesPage(first, pageSize int) types.NameSummary {
	copy := a.namesSum
	first = base.Max(0, base.Min(first, len(copy.Names)-1))
	last := base.Min(len(copy.Names), first+pageSize)
	copy.Names = copy.Names[first:last]
	return copy
}

func (a *App) GetNamesCnt() int {
	return len(a.namesSum.Names)
}

func (a *App) loadNames() error {
	nameTypes := []coreTypes.Parts{coreTypes.Regular, coreTypes.Custom, coreTypes.Prefund, coreTypes.Baddress}
	for _, t := range nameTypes {
		if namesMap, err := names.LoadNamesMap("mainnet", t, nil); err != nil {
			return err
		} else {
			for addr, name := range namesMap {
				name.Parts |= t
				if vv, ok := a.namesSum.NamesMap[addr]; ok {
					name = vv
					name.Parts |= t
				}
				a.namesSum.NamesMap[addr] = name
			}
		}
	}
	for _, name := range a.namesSum.NamesMap {
		a.namesSum.Names = append(a.namesSum.Names, name)
	}
	sort.Slice(a.namesSum.Names, func(i, j int) bool {
		ti := a.namesSum.Names[i].Parts
		if ti == coreTypes.Regular {
			ti = 7
		}
		tj := a.namesSum.Names[j].Parts
		if tj == coreTypes.Regular {
			tj = 7
		}
		if ti == tj {
			if a.namesSum.Names[i].Tags == a.namesSum.Names[j].Tags {
				return a.namesSum.Names[i].Address.Hex() < a.namesSum.Names[j].Address.Hex()
			}
			return a.namesSum.Names[i].Tags < a.namesSum.Names[j].Tags
		}
		return ti < tj
	})
	a.namesSum.Summarize()
	return nil
}
