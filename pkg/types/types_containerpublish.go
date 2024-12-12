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

type PublishContainer struct {
	Chain   string       `json:"chain"`
	Items   []CacheItem  `json:"items"`
	NItems  uint64       `json:"nItems"`
	Updater sdk.Updater  `json:"updater"`
	Sorts   sdk.SortSpec `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewPublishContainer(chain string, itemsIn []CacheItem) PublishContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := PublishContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewPublishUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewPublishUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Path: walk.GetRootPathFromCacheType(chain, walk.Index_Bloom), Type: sdk.Folder},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("publish", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *PublishContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *PublishContainer) GetItems() interface{} {
	return s.Items
}

func (s *PublishContainer) SetItems(items interface{}) {
	s.Items = items.([]CacheItem)
}

func (s *PublishContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *PublishContainer) ShallowCopy() Containerer {
	ret := &PublishContainer{
		Chain:   s.Chain,
		NItems:  s.NItems,
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *PublishContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *PublishContainer) passesFilter(item *CacheItem, filter *Filter) (ret bool) {
	_ = item // linter
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *PublishContainer) Accumulate(item *CacheItem) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *PublishContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *PublishContainer) CollateAndFilter(filter *Filter) interface{} {
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

func (s *PublishContainer) ForEveryItem(process EveryCacheItemFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *PublishContainer) Sort() (err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
