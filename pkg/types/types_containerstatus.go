// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type StatusContainer struct {
	NBytes   uint64 `json:"nBytes"`
	NFiles   uint64 `json:"nFiles"`
	NFolders uint64 `json:"nFolders"`
	Status   `json:",inline"`
	Updater  sdk.Updater `json:"updater"`
	Items    []CacheItem `json:"items"`
	NItems   uint64      `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewStatusContainer(chain string, itemsIn []CacheItem, status *Status) StatusContainer {
	ret := StatusContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Status:  *status,
		Updater: NewStatusUpdater(chain),
	}
	// EXISTING_CODE
	ret.Items = status.Caches
	ret.NItems = uint64(len(ret.Items))
	// EXISTING_CODE
	return ret
}

func NewStatusUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Duration: 2 * time.Minute, Type: sdk.Timer},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("status", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *StatusContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *StatusContainer) GetItems() interface{} {
	return s.Items
}

func (s *StatusContainer) SetItems(items interface{}) {
	s.Items = items.([]CacheItem)
}

func (s *StatusContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *StatusContainer) ShallowCopy() Containerer {
	ret := &StatusContainer{
		NBytes:   s.NBytes,
		NFiles:   s.NFiles,
		NFolders: s.NFolders,
		Status:   s.Status.ShallowCopy(),
		Updater:  s.Updater,
		NItems:   s.NItems,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *StatusContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	s.NFolders = 0
	s.NFiles = 0
	s.NBytes = 0
	// EXISTING_CODE
}

func (s *StatusContainer) passesFilter(item *CacheItem, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *StatusContainer) Accumulate(item *CacheItem) {
	s.NItems++
	// EXISTING_CODE
	s.NFolders += item.NFolders
	s.NFiles += item.NFiles
	s.NBytes += uint64(item.SizeInBytes)
	// EXISTING_CODE
}

func (s *StatusContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *StatusContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("status") // may be empty
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
		if s.passesFilter(item, &filter) {
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

func (s *StatusContainer) ForEveryItem(process EveryCacheItemFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE
