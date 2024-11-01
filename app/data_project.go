// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

// var projectLock atomic.Uint32

func (a *App) ProjectPage(first, pageSize int) *types.ProjectContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	// first = base.Max(0, base.Min(first, len(a.projects.Items)-1))
	// last := base.Min(len(a.projects.Items), first+pageSize)
	// copy, _ := a.projects.ShallowCopy().(*types.ProjectContainer)
	// copy.Items = a.projects.Items[first:last]
	// return copy
	return &types.ProjectContainer{}
}

// func (a *App) loadProjects(wg *sync.WaitGroup, errorChan chan error) error {
// 	defer func() {
// 		if wg != nil {
// 			wg.Done()
// 		}
// 	}()

// 	if !projectLock.CompareAndSwap(0, 1) {
// 		return nil
// 	}
// 	defer projectLock.CompareAndSwap(1, 0)

// 	if !a.projects.NeedsUpdate(a.forceProject()) {
// 		return nil
// 	}

// 	// EXISTING_CODE
// 	items := []base.Address{}
// 	a.HistoryCache.Range(func(address base.Address, _ types.HistoryContainer) bool {
// 		items = append(items, address)
// 		return true
// 	})
// 	// EXISTING_CODE

// 	{

// 		// EXISTING_CODE
// 		// EXISTING_CODE
// 		// EXISTING_CODE
// 		// EXISTING_CODE
// 	}

// 	a.projects = types.NewProjectContainer(a.session.LastChain, items)
// 	sort.Slice(a.projects.Items, func(i, j int) bool {
// 		return a.projects.Items[i].Address.Cmp(a.projects.Items[j].Address) < 0
// 	})
// 	messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{
// 		String1: "Project " + fmt.Sprintf("%d", len(a.projects.Items)),
// 	})

// 	return nil
// }

// func (a *App) forceProject() (force bool) {
// 	// EXISTING_CODE
// 	force = a.forceName()
// 	// EXISTING_CODE
// 	return
// }

// EXISTING_CODE
func (a *App) ModifyProject(modData *ModifyData) {
	a.cancelContext(modData.Address)
	a.HistoryCache.Delete(modData.Address)
	for _, project := range a.dashboard.Projects {
		// x
		for i, history := range project.Histories {
			if history == modData.Address {
				a.dashboard.Projects = append(a.dashboard.Projects[:i], a.dashboard.Projects[i+1:]...)
				break
			}
		}
	}
	a.loadDashboard(nil, nil)
}

func (a *App) Reload(address base.Address) {
	a.ModifyProject(&ModifyData{
		Operation: "reload",
		Address:   address,
	})
	a.loadHistory(a.GetAddress(), nil, nil)
	_ = a.Refresh()
	a.loadDashboard(nil, nil)
}

func (a *App) GoToHistory(address base.Address) {
	a.SetRoute("/history", address.Hex())
	a.Reload(address)

	route := "/history/" + address.Hex()
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: route,
	})
}

// EXISTING_CODE
