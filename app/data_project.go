// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
)

// EXISTING_CODE

var projectLock atomic.Uint32

func (a *App) ProjectPage(first, pageSize int) *types.ProjectContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.projects.Items)-1))
	last := base.Min(len(a.projects.Items), first+pageSize)
	copy, _ := a.projects.ShallowCopy().(*types.ProjectContainer)
	copy.Items = a.projects.Items[first:last]
	return copy
}

func (a *App) loadProjects(wg *sync.WaitGroup, errorChan chan error) error {
	_ = errorChan
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !projectLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer projectLock.CompareAndSwap(1, 0)

	// EXISTING_CODE
	items := []types.HistoryContainer{}
	a.historyCache.Range(func(_ base.Address, h types.HistoryContainer) bool {
		items = append(items, h)
		return true
	})
	sort.Slice(items, func(i, j int) bool {
		return items[i].Address.Hex() < items[j].Address.Hex()
	})
	a.projects = types.NewProjectContainer(a.session.LastChain, items)
	if !a.projects.NeedsUpdate(a.forceProject()) {
		return nil
	}

	// EXISTING_CODE

	{

		// EXISTING_CODE
		// EXISTING_CODE
		// EXISTING_CODE
		// EXISTING_CODE
	}

	a.projects.NItems = uint64(len(a.projects.Items))
	a.projects.NMonitors = uint64(len(a.monitors.Items))
	a.projects.NNames = uint64(len(a.names.Items))
	a.projects.NAbis = uint64(len(a.abis.Items))
	a.projects.NIndexes = uint64(len(a.indexes.Items))
	a.projects.NManifests = uint64(len(a.manifests.Items))
	a.projects.NCaches = uint64(len(a.status.Caches))
	a.projects.ForEveryHistory(func(item *types.HistoryContainer, data any) bool {
		item.Summarize()
		a.projects.HistorySize += uint64(item.SizeOf())
		return true
	}, nil)
	sort.Slice(a.projects.Items, func(i, j int) bool {
		ai := a.projects.Items[i].Address
		aj := a.projects.Items[j].Address
		return ai.Hex() < aj.Hex()
	})
	a.emitInfoMsg("Loaded projects", "")
	return nil
}

func (a *App) forceProject() (force bool) {
	// EXISTING_CODE
	force = a.forceName()
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (a *App) ModifyHistory(modData *ModifyData) {
	switch crud.OpFromString(modData.Operation) {
	case crud.Delete:
		a.historyCache.Delete(modData.Address)
		a.GoToAddress(modData.Address)
		a.dirty = true
		a.emitInfoMsg(a.getFullPath(), fmt.Sprint("deleted address", modData.Address.Hex()))
	}
}

// EXISTING_CODE
