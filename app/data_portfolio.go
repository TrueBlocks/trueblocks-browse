package app

import (
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) PortfolioPage(first, pageSize int) *types.PortfolioContainer {
	first = base.Max(0, base.Min(first, len(a.portfolio.Items)-1))
	last := base.Min(len(a.portfolio.Items), first+pageSize)
	// copy, _ := a.portfolio.ShallowCopy().(*types.PortfolioContainer)
	copy := a.portfolio.ShallowCopy()
	copy.Items = a.portfolio.Items[first:last]
	return &copy
}

var portfolioLock atomic.Uint32

func (a *App) loadPortfolio(wg *sync.WaitGroup, errorChan chan error) error {
	if !portfolioLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer portfolioLock.CompareAndSwap(1, 0)

	_ = errorChan // delint
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	containers := []types.Containerer{
		&a.abis,
		// &HistoryContainer{},
		&a.index,
		&a.manifest,
		&a.monitors,
		&a.names,
		&a.status,
		// &PortfolioContainer{},
	}
	needsUpdate := false
	for _, container := range containers {
		if container.NeedsUpdate() {
			needsUpdate = true
			break
		}
	}
	if !needsUpdate && len(a.historyMap) == a.portfolio.MyCount {
		return nil
	}

	a.portfolio = types.PortfolioContainer{}
	a.portfolio.MyCount = len(a.historyMap)
	a.portfolio.NMonitors = len(a.monitors.Items)
	a.portfolio.NNames = len(a.names.Names)
	a.portfolio.NAbis = len(a.abis.Items)
	a.portfolio.NIndexes = len(a.index.Items)
	a.portfolio.NManifests = len(a.manifest.Items)
	a.portfolio.NCaches = len(a.status.Items)
	a.portfolio.HistorySize = 0
	for _, m := range a.historyMap {
		a.portfolio.Summary.Balance += m.Balance
		a.portfolio.Summary.NItems += m.NItems
		a.portfolio.Summary.NLogs += m.NLogs
		a.portfolio.Summary.NErrors += m.NErrors
		a.portfolio.Summary.NTokens += m.NTokens
		m.Summarize()
		a.portfolio.Items = append(a.portfolio.Items, m.ShallowCopy())
		a.portfolio.HistorySize += m.SizeOf()
	}
	sort.Slice(a.portfolio.Items, func(i, j int) bool {
		return a.portfolio.Items[i].Address.Cmp(a.portfolio.Items[j].Address.Address) < 0
	})
	messages.SendInfo(a.ctx, "Loaded portfolio")

	return nil
}
