// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"unsafe"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type HistoryContainer struct {
	Address base.Address            `json:"address"`
	Balance string                  `json:"balance"`
	Chain   string                  `json:"chain"`
	NErrors uint64                  `json:"nErrors"`
	NLogs   uint64                  `json:"nLogs"`
	NTokens uint64                  `json:"nTokens"`
	NTotal  uint64                  `json:"nTotal"`
	Name    string                  `json:"name"`
	Updater updater.Updater         `json:"updater"`
	Items   []coreTypes.Transaction `json:"items"`
	NItems  uint64                  `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewHistoryContainer(chain string, itemsIn []coreTypes.Transaction, address base.Address) HistoryContainer {
	ret := HistoryContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Address: address,
		Chain:   chain,
		Updater: NewHistoryUpdater(chain, address),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewHistoryUpdater(chain string, address base.Address, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: coreMonitor.PathToMonitorFile(chain, address), Type: updater.FileSize},
		{Path: coreConfig.MustGetPathToChainConfig(namesChain), Type: updater.Folder},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("history", items)
	if reset {
		updater.Reset()
	}
	return updater
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
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *HistoryContainer) ShallowCopy() Containerer {
	ret := &HistoryContainer{
		Address: s.Address,
		Balance: s.Balance,
		Chain:   s.Chain,
		NErrors: s.NErrors,
		NLogs:   s.NLogs,
		NTokens: s.NTokens,
		NTotal:  s.NTotal,
		Name:    s.Name,
		Updater: s.Updater,
		NItems:  s.NItems,
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
