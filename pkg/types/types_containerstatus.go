// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type StatusContainer struct {
	NBytes           uint64 `json:"nBytes"`
	NFiles           uint64 `json:"nFiles"`
	NFolders         uint64 `json:"nFolders"`
	coreTypes.Status `json:",inline"`
	Items            []coreTypes.CacheItem `json:"items"`
	NItems           uint64                `json:"nItems"`
	LastUpdate       int64                 `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewStatusContainer(chain string, itemsIn []coreTypes.CacheItem, status *coreTypes.Status) StatusContainer {
	ret := StatusContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Status: *status,
	}
	ret.Chain = chain
	ret.LastUpdate, _ = ret.getStatusReload()
	// EXISTING_CODE
	ret.LastUpdate = time.Now().Unix()
	ret.Items = status.Caches
	ret.NItems = uint64(len(ret.Items))
	// EXISTING_CODE
	return ret
}

func (s *StatusContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *StatusContainer) GetItems() interface{} {
	return s.Items
}

func (s *StatusContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.CacheItem)
}

func (s *StatusContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getStatusReload()
	if force || reload {
		DebugInts("status", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *StatusContainer) ShallowCopy() Containerer {
	ret := &StatusContainer{
		NBytes:     s.NBytes,
		NFiles:     s.NFiles,
		NFolders:   s.NFolders,
		Status:     s.Status.ShallowCopy(),
		NItems:     s.NItems,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	ret.Chain = s.Chain
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

func (s *StatusContainer) passesFilter(item *coreTypes.CacheItem, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *StatusContainer) Accumulate(item *coreTypes.CacheItem) {
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
	filtered := []coreTypes.CacheItem{}
	s.ForEveryItem(func(item *coreTypes.CacheItem, data any) bool {
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

func (s *StatusContainer) getStatusReload() (ret int64, reload bool) {
	// EXISTING_CODE
	ret = time.Now().Unix()
	reload = ret > s.LastUpdate+(60*2) // every two minutes
	// EXISTING_CODE
	return
}

type EveryCacheItemFn func(item *coreTypes.CacheItem, data any) bool

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
