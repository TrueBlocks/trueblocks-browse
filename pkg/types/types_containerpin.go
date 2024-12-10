// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type PinContainer struct {
	Chain   string       `json:"chain"`
	Items   []CacheItem  `json:"items"`
	NItems  uint64       `json:"nItems"`
	Updater sdk.Updater  `json:"updater"`
	Sorts   sdk.SortSpec `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewPinContainer(chain string, itemsIn []CacheItem) PinContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := PinContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewPinUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewPinUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Path: walk.GetRootPathFromCacheType(chain, walk.Index_Bloom), Type: sdk.Folder},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("pins", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *PinContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *PinContainer) GetItems() interface{} {
	return s.Items
}

func (s *PinContainer) SetItems(items interface{}) {
	s.Items = items.([]CacheItem)
}

func (s *PinContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *PinContainer) ShallowCopy() Containerer {
	ret := &PinContainer{
		Chain:   s.Chain,
		NItems:  s.NItems,
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *PinContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *PinContainer) passesFilter(item *CacheItem, filter *Filter) (ret bool) {
	_ = item // linter
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *PinContainer) Accumulate(item *CacheItem) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *PinContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *PinContainer) CollateAndFilter(filter *Filter) interface{} {
	s.Clear()

	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *CacheItem, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []CacheItem{}
	s.ForEveryItem(func(item *CacheItem, data any) bool {
		if s.passesFilter(item, filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

func (s *PinContainer) ForEveryItem(process EveryCacheItemFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *PinContainer) Sort() (err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
