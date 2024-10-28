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
	copy, _ := a.project.ShallowCopy().(*types.ProjectContainer)
	copy.Items = a.project.Items[first:last]
	return copy
}

var projectLock atomic.Uint32

func (a *App) loadProject(wg *sync.WaitGroup, errorChan chan error) error {
	_ = errorChan // delint

	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !projectLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer projectLock.CompareAndSwap(1, 0)

	// containers := []types.Containerer{
	// 	&a.abis,
	// 	// &HistoryContainer{},
	// 	&a.indexes,
	// 	&a.manifests,
	// 	&a.monitors,
	// 	&a.names,
	// 	&a.status,
	// 	// &ProjectContainer{},
	// }
	// needsUpdate := false
	// for _, container := range containers {
	// 	if container.NeedsUpdate(a.forceProject()) {
	// 		needsUpdate = true
	// 		break
	// 	}
	// }
	// if !needsUpdate && a.openFileCnt() == a.project.NOpenFiles {
	// 	return nil
	// }
	_ = a.forceProject() // silence unused

	a.project = types.NewProjectContainer(
		a.project.Filename,
		a.project.HistoryMap,
		a.project.BalanceMap,
		a.project.EnsMap,
	)
	a.project.NItems = uint64(a.openFileCnt())
	a.project.NMonitors = uint64(len(a.monitors.Items))
	a.project.NNames = uint64(len(a.names.Items))
	a.project.NAbis = uint64(len(a.abis.Items))
	a.project.NIndexes = uint64(len(a.indexes.Items))
	a.project.NManifests = uint64(len(a.manifests.Items))
	a.project.NCaches = uint64(len(a.status.Caches))
	_ = a.forEveryHistory(func(item *types.HistoryContainer) bool {
		a.project.Summary.Balance += item.Balance
		a.project.Summary.NItems += item.NItems
		a.project.Summary.NTotal += item.NTotal
		a.project.Summary.NLogs += item.NLogs
		a.project.Summary.NErrors += item.NErrors
		a.project.Summary.NTokens += item.NTokens
		item.Summarize()
		if copy, ok := item.ShallowCopy().(*types.HistoryContainer); ok {
			a.project.Items = append(a.project.Items, *copy)
		}
		a.project.HistorySize += uint64(item.SizeOf())
		return true
	})
	sort.Slice(a.project.Items, func(i, j int) bool {
		return a.project.Items[i].Address.Cmp(a.project.Items[j].Address.Address) < 0
	})

	return nil
}

func (a *App) ModifyProject(modData *ModifyData) {
	a.CancelContext(modData.Address)
	a.project.HistoryMap.Delete(modData.Address)
	for i, item := range a.project.Items {
		if item.Address == modData.Address {
			a.project.Items = append(a.project.Items[:i], a.project.Items[i+1:]...)
			break
		}
	}
	a.loadProject(nil, nil)
}

func (a *App) forceProject() bool {
	return a.forceName()
}

func (a *App) Reload(address base.Address) {
	a.ModifyProject(&ModifyData{
		Operation: "reload",
		Address:   address,
	})
	a.loadHistory(a.GetAddress(), nil, nil)
	_ = a.Refresh()
	a.loadProject(nil, nil)
}

func (a *App) GoToHistory(address base.Address) {
	a.SetRoute("/history", address.Hex())
	a.Reload(address)

	route := "/history/" + address.Hex()
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: route,
	})
}
