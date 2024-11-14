// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

func (a *App) FetchName(first, pageSize int) *types.NameContainer {
	// EXISTING_CODE
	namesMutex.Lock()
	defer namesMutex.Unlock()
	// EXISTING_CODE

	a.names.CollateAndFilter()
	filtered := []coreTypes.Name{}
	filterStr := ""
	if filter, exists := a.filterMap.Load("names"); exists {
		filterStr = filter.(Filter).Criteria
	}
	if len(filterStr) > 0 {
		for _, item := range a.names.Items {
			s := strings.ToLower(filterStr)
			n := strings.ToLower(item.Name)
			a := strings.ToLower(item.Address.Hex())
			t := strings.ToLower(item.Tags)
			c1 := strings.Contains(n, s)
			c2 := strings.Contains(a, s)
			c3 := strings.Contains(t, s)
			if c1 || c2 || c3 {
				filtered = append(filtered, item)
			}
		}
	} else {
		filtered = a.names.Items
	}

	first = base.Max(0, base.Min(first, len(filtered)-1))
	last := base.Min(len(filtered), first+pageSize)
	copy, _ := a.names.ShallowCopy().(*types.NameContainer)
	copy.Items = filtered[first:last]
	return copy
}
