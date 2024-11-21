// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"unsafe"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type HistoryContainer struct {
	Address    base.Address            `json:"address"`
	Balance    string                  `json:"balance"`
	Chain      string                  `json:"chain"`
	LastUpdate int64                   `json:"lastUpdate"`
	NErrors    uint64                  `json:"nErrors"`
	NLogs      uint64                  `json:"nLogs"`
	NTokens    uint64                  `json:"nTokens"`
	NTotal     uint64                  `json:"nTotal"`
	Name       string                  `json:"name"`
	Items      []coreTypes.Transaction `json:"items"`
	NItems     uint64                  `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewHistoryContainer(chain string, itemsIn []coreTypes.Transaction, address base.Address) HistoryContainer {
	ret := HistoryContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
	}
	ret.Chain = chain
	ret.LastUpdate, _ = ret.getHistoryReload()
	// EXISTING_CODE
	ret.Address = address
	ret.LastUpdate, _ = ret.getHistoryReload() // DO NOT REMOVE (needs address)
	// EXISTING_CODE
	return ret
}

func (s *HistoryContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *HistoryContainer) GetItems() interface{} {
	return s.Items
}

func (s *HistoryContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.Transaction)
}

func (s *HistoryContainer) NeedsUpdate() bool {
	latest, reload := s.getHistoryReload()
	if reload {
		DebugInts("history", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *HistoryContainer) ShallowCopy() Containerer {
	ret := &HistoryContainer{
		Address:    s.Address,
		Balance:    s.Balance,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		NErrors:    s.NErrors,
		NLogs:      s.NLogs,
		NTokens:    s.NTokens,
		NTotal:     s.NTotal,
		Name:       s.Name,
		NItems:     s.NItems,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *HistoryContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	s.NLogs = 0
	s.NTokens = 0
	s.NErrors = 0
	// EXISTING_CODE
}

func (s *HistoryContainer) passesFilter(item *coreTypes.Transaction, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *HistoryContainer) Accumulate(item *coreTypes.Transaction) {
	s.NItems++
	// EXISTING_CODE
	if item.Receipt != nil {
		s.NLogs += uint64(len(item.Receipt.Logs))
	}
	if item.HasToken {
		s.NTokens++
	}
	if item.IsError {
		s.NErrors++
	}
	// EXISTING_CODE
}

func (s *HistoryContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *HistoryContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("history") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *coreTypes.Transaction, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []coreTypes.Transaction{}
	s.ForEveryItem(func(item *coreTypes.Transaction, data any) bool {
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

func (s *HistoryContainer) getHistoryReload() (ret int64, reload bool) {
	// EXISTING_CODE
	fn := coreMonitor.PathToMonitorFile(s.Chain, s.Address)
	ret = file.FileSize(fn)
	reload = ret > s.LastUpdate
	// EXISTING_CODE
	return
}

type EveryTransactionFn func(item *coreTypes.Transaction, data any) bool

func (s *HistoryContainer) ForEveryItem(process EveryTransactionFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
func (s *HistoryContainer) SizeOf() int {
	size := unsafe.Sizeof(s.Address) + unsafe.Sizeof(s.Name) + unsafe.Sizeof(s.Balance) + unsafe.Sizeof(s.NLogs) + unsafe.Sizeof(s.NTokens) + unsafe.Sizeof(s.NErrors) + unsafe.Sizeof(s.NItems) + unsafe.Sizeof(s.NTotal)
	for _, record := range s.Items {
		size += unsafe.Sizeof(record)
	}
	return int(size)
}

// EXISTING_CODE
