package app

import (
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) ProjectPage(first, pageSize int) *types.ProjectContainer {
	first = base.Max(0, base.Min(first, len(a.project.Items)-1))
	last := base.Min(len(a.project.Items), first+pageSize)
	// copy, _ := a.project.ShallowCopy().(*types.ProjectContainer)
	copy := a.project.ShallowCopy()
	copy.Items = a.project.Items[first:last]
	return &copy
}

var projectLock atomic.Uint32

func (a *App) loadProject(wg *sync.WaitGroup, errorChan chan error) error {
	if !projectLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer projectLock.CompareAndSwap(1, 0)

	_ = errorChan // delint
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	// containers := []types.Containerer{
	// 	&a.abis,
	// 	// &HistoryContainer{},
	// 	&a.index,
	// 	&a.manifest,
	// 	&a.monitors,
	// 	&a.names,
	// 	&a.status,
	// 	// &ProjectContainer{},
	// }
	// needsUpdate := false
	// for _, container := range containers {
	// 	if container.NeedsUpdate() {
	// 		needsUpdate = true
	// 		break
	// 	}
	// }
	// if !needsUpdate && a.OpenFileCnt() == a.project.NOpenFiles {
	// 	return nil
	// }

	a.project = types.ProjectContainer{}
	a.project.NOpenFiles = a.OpenFileCnt()
	a.project.NMonitors = len(a.monitors.Items)
	a.project.NNames = len(a.names.Names)
	a.project.NAbis = len(a.abis.Items)
	a.project.NIndexes = len(a.index.Items)
	a.project.NManifests = len(a.manifest.Items)
	a.project.NCaches = len(a.status.Items)
	a.project.HistorySize = 0
	for _, m := range a.historyMap {
		a.project.Summary.Balance += m.Balance
		a.project.Summary.NItems += m.NItems
		a.project.Summary.NLogs += m.NLogs
		a.project.Summary.NErrors += m.NErrors
		a.project.Summary.NTokens += m.NTokens
		m.Summarize()
		if copy, ok := m.ShallowCopy().(*types.HistoryContainer); ok {
			a.project.Items = append(a.project.Items, *copy)
		}
		a.project.HistorySize += m.SizeOf()
	}
	sort.Slice(a.project.Items, func(i, j int) bool {
		return a.project.Items[i].Address.Cmp(a.project.Items[j].Address.Address) < 0
	})
	messages.SendInfo(a.ctx, "Loaded project")

	return nil
}
