// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

var dashboardLock atomic.Uint32

func (a *App) DashboardPage(first, pageSize int) *types.DashboardContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	first = base.Max(0, base.Min(first, len(a.dashboard.Projects)-1))
	last := base.Min(len(a.dashboard.Projects), first+pageSize)
	copy, _ := a.dashboard.ShallowCopy().(*types.DashboardContainer)
	copy.Projects = a.dashboard.Projects[first:last]
	return copy
}

func (a *App) loadDashboard(wg *sync.WaitGroup, errorChan chan error) error {
	_ = errorChan
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !dashboardLock.CompareAndSwap(0, 1) {
		return nil
	}
	defer dashboardLock.CompareAndSwap(1, 0)

	// if !a.dashboard.NeedsUpdate(a.forceDashboard()) {
	// 	return nil
	// }

	// opts := sdk.dashboardOptions{
	// 	Globals: a.toGlobals(),
	// }
	// EXISTING_CODE
	// EXISTING_CODE
	// opts.Verbose = true

	// if dashboard, meta, err := opts.dashboardList(); err != nil {
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else if (dashboard == nil) || (len(dashboard) == 0) {
	// 	err = fmt.Errorf("no dashboard found")
	// 	if errorChan != nil {
	// 		errorChan <- err
	// 	}
	// 	return err
	// } else {
	// EXISTING_CODE
	// EXISTING_CODE
	// a.meta = *meta
	projects := []types.ProjectContainer{}
	a.ProjectCache.ForEveryProject(func(_ string, p types.ProjectContainer) bool {
		projects = append(projects, p)
		return true
	})
	a.dashboard = types.NewDashboardContainer(a.Chain, projects)
	a.dashboard.NProjects = uint64(len(a.dashboard.Projects))
	// a.dashboard.NMonitors = uint64(len(a.monitors.Items))
	// a.dashboard.NNames = uint64(len(a.names.Items))
	// a.dashboard.NAbis = uint64(len(a.abis.Items))
	// a.dashboard.NIndexes = uint64(len(a.indexes.Items))
	// a.dashboard.NManifests = uint64(len(a.manifests.Items))
	// a.dashboard.NCaches = uint64(len(a.status.Caches))
	// _ = a.forEveryHistory(func(item *types.HistoryContainer) bool {
	// 	item.Summarize()
	// 	a.dashboard.HistorySize += uint64(item.SizeOf())
	// 	return true
	// })
	// EXISTING_CODE
	// EXISTING_CODE
	a.dashboard.Summarize()
	messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{String1: "Loaded dashboard"})
	// }

	return nil
}

func (a *App) forceDashboard() (force bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
