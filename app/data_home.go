package app

import (
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type HomeContainer struct {
	Summary    types.HistoryContainer   `json:",inline"`
	Items      []types.HistoryContainer `json:"items"`
	NMonitors  int                      `json:"nMonitors"`
	NNames     int                      `json:"nNames"`
	NAbis      int                      `json:"nAbis"`
	NIndexes   int                      `json:"nIndexes"`
	NManifests int                      `json:"nManifests"`
	NCaches    int                      `json:"nCaches"`
}

func (h *HomeContainer) ShallowCopy() HomeContainer {
	ret := HomeContainer{}
	ret.Summary = h.Summary.ShallowCopy()
	// ret.Items = h.Items
	ret.NMonitors = h.NMonitors
	ret.NNames = h.NNames
	ret.NAbis = h.NAbis
	ret.NIndexes = h.NIndexes
	ret.NManifests = h.NManifests
	ret.NCaches = h.NCaches
	return ret
}

func (a *App) HomePage(first, pageSize int) HomeContainer {
	first = base.Max(0, base.Min(first, len(a.home.Items)-1))
	last := base.Min(len(a.home.Items), first+pageSize)
	copy := a.home.ShallowCopy()
	copy.Items = a.home.Items[first:last]
	return copy
}

func (a *App) loadHome(wg *sync.WaitGroup, errorChan chan error) error {
	a.home = HomeContainer{}
	a.home.NMonitors = len(a.monitors.Items)
	a.home.NNames = len(a.names.Names)
	a.home.NAbis = len(a.abis.Items)
	a.home.NIndexes = len(a.index.Items)
	a.home.NManifests = len(a.manifest.Items)
	a.home.NCaches = len(a.status.Items)
	for _, m := range a.historyMap {
		a.home.Summary.NItems += m.NItems
		a.home.Summary.NLogs += m.NLogs
		a.home.Summary.NErrors += m.NErrors
		a.home.Summary.NTokens += m.NTokens
		a.home.Items = append(a.home.Items, m.ShallowCopy())
	}
	sort.Slice(a.home.Items, func(i, j int) bool {
		return a.home.Items[i].Name < a.home.Items[j].Name
	})

	return nil
}
