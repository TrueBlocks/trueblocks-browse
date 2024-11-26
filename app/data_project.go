// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"sort"
	"sync"
	"sync/atomic"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

var projectLock atomic.Uint32

func (a *App) loadProject(wg *sync.WaitGroup, errorChan chan error) error {
	defer a.trackPerformance("loadProject", false)()
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
	// EXISTING_CODE

	if !a.project.NeedsUpdate() {
		return nil
	}
	updater := a.project.Updater
	defer func() {
		a.project.Updater = updater
	}()
	logger.InfoBY("Updating needed for project...")

	opts := ProjectOptions{
		Globals: a.getGlobals(true /* verbose */),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	if items, meta, err := opts.ProjectList(); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (items == nil) || (len(items) == 0) {
		// expected outcome
		a.meta = *meta
		return nil
	} else {
		// EXISTING_CODE
		items := []types.HistoryContainer{}
		// HIST-PROJ
		a.historyCache.Range(func(_ base.Address, h types.HistoryContainer) bool {
			items = append(items, h)
			return true
		})
		// EXISTING_CODE
		a.meta = *meta
		a.project = types.NewProjectContainer(opts.Chain, items)
		// EXISTING_CODE
		a.project.NItems = uint64(len(a.project.Items))
		a.project.NMonitors = uint64(len(a.monitors.Items))
		a.project.NNames = uint64(len(a.names.Items))
		a.project.NAbis = uint64(len(a.abis.Items))
		a.project.NIndexes = uint64(len(a.indexes.Items))
		a.project.NManifests = uint64(len(a.manifests.Items))
		a.project.NCaches = uint64(len(a.status.Caches))
		a.project.ForEveryItem(func(item *types.HistoryContainer, data any) bool {
			a.project.HistorySize += uint64(item.SizeOf())
			return true
		}, nil)
		sort.Slice(a.project.Items, func(i, j int) bool {
			return a.project.Items[i].Address.Hex() < a.project.Items[j].Address.Hex()
		})
		// EXISTING_CODE
		a.emitLoadingMsg(messages.Loaded, "project")
	}

	return nil
}

// EXISTING_CODE
type ProjectOptions struct {
	Globals sdk.Globals
	Chain   string
}

func (opts *ProjectOptions) ProjectList() ([]types.HistoryContainer, *types.Meta, error) {
	meta, err := sdk.GetMetaData(namesChain)
	return []types.HistoryContainer{}, meta, err
}

// EXISTING_CODE
