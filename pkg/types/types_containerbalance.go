// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"unsafe"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type BalanceContainer struct {
	Address base.Address  `json:"address"`
	Balance string        `json:"balance"`
	Chain   string        `json:"chain"`
	Items   []Transaction `json:"items"`
	NErrors uint64        `json:"nErrors"`
	NItems  uint64        `json:"nItems"`
	NLogs   uint64        `json:"nLogs"`
	NTokens uint64        `json:"nTokens"`
	NTotal  uint64        `json:"nTotal"`
	Name    string        `json:"name"`
	Updater sdk.Updater   `json:"updater"`
	Sorts   sdk.SortSpec  `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewBalanceContainer(chain string, itemsIn []Transaction, address base.Address) BalanceContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := BalanceContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewBalanceUpdater(chain, address),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewBalanceUpdater(chain string, address base.Address, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Path: coreMonitor.PathToMonitorFile(chain, address), Type: sdk.FileSize},
		{Path: coreConfig.MustGetPathToChainConfig(namesChain), Type: sdk.Folder},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("balances", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *BalanceContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *BalanceContainer) GetItems() interface{} {
	return s.Items
}

func (s *BalanceContainer) SetItems(items interface{}) {
	s.Items = items.([]Transaction)
}

func (s *BalanceContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *BalanceContainer) ShallowCopy() Containerer {
	ret := &BalanceContainer{
		Address: s.Address,
		Balance: s.Balance,
		Chain:   s.Chain,
		NErrors: s.NErrors,
		NItems:  s.NItems,
		NLogs:   s.NLogs,
		NTokens: s.NTokens,
		NTotal:  s.NTotal,
		Name:    s.Name,
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *BalanceContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *BalanceContainer) passesFilter(item *Transaction, filter *Filter) (ret bool) {
	_ = item // linter
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *BalanceContainer) Accumulate(item *Transaction) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *BalanceContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *BalanceContainer) CollateAndFilter(filter *Filter) interface{} {
	s.Clear()

	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *Transaction, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []Transaction{}
	s.ForEveryItem(func(item *Transaction, data any) bool {
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

func (s *BalanceContainer) ForEveryItem(process EveryTransactionFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *BalanceContainer) Sort() (err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
/*
type HistoryContainer struct {
	Address base.Address  `json:"address"`
	Balance string        `json:"balance"`
	Chain   string        `json:"chain"`
	Items   []Transaction `json:"items"`
	NErrors uint64        `json:"nErrors"`
	NItems  uint64        `json:"nItems"`
	NLogs   uint64        `json:"nLogs"`
	NTokens uint64        `json:"nTokens"`
	NTotal  uint64        `json:"nTotal"`
	Name    string        `json:"name"`
	Updater sdk.Updater   `json:"updater"`
	Sorts   sdk.SortSpec  `json:"sorts"`
}

func NewHistoryContainer(chain string, itemsIn []Transaction, address base.Address) HistoryContainer {
	ret := HistoryContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewHistoryUpdater(chain, address),
	}
	ret.Chain = chain
	ret.Address = address
	return ret
}

func NewHistoryUpdater(chain string, address base.Address, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	items := []sdk.UpdaterItem{
		{Path: coreMonitor.PathToMonitorFile(chain, address), Type: sdk.FileSize},
		{Path: coreConfig.MustGetPathToChainConfig(namesChain), Type: sdk.Folder},
	}
	u, _ := sdk.NewUpdater("history", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *HistoryContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *HistoryContainer) GetItems() interface{} {
	return s.Items
}

func (s *HistoryContainer) SetItems(items interface{}) {
	s.Items = items.([]Transaction)
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
		NItems:  s.NItems,
		NLogs:   s.NLogs,
		NTokens: s.NTokens,
		NTotal:  s.NTotal,
		Name:    s.Name,
		Updater: s.Updater,
	}
	return ret
}

func (s *HistoryContainer) Clear() {
	s.NItems = 0
	s.NLogs = 0
	s.NTokens = 0
	s.NErrors = 0
}

func (s *HistoryContainer) passesFilter(item *Transaction, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
	}
	return
}

func (s *HistoryContainer) Accumulate(item *Transaction) {
	s.NItems++
	if item.Receipt != nil {
		s.NLogs += uint64(len(item.Receipt.Logs))
	}
	if item.HasToken {
		s.NTokens++
	}
	if item.IsError {
		s.NErrors++
	}
}

func (s *HistoryContainer) Finalize() {
}

func (s *HistoryContainer) CollateAndFilter(filter *Filter) interface{} {
	s.Clear()

	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *Transaction, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []Transaction{}
	s.ForEveryItem(func(item *Transaction, data any) bool {
		if s.passesFilter(item, filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	return filtered
}

func (s *HistoryContainer) ForEveryItem(process EveryTransactionFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *HistoryContainer) Sort() (err error) {
	// TODO: Is there anything to sort for history?
	return
}

*/
func (s *HistoryContainer) ForEveryItem(process EveryTransactionFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *HistoryContainer) SizeOf() int {
	size := unsafe.Sizeof(s.Address) + unsafe.Sizeof(s.Name) + unsafe.Sizeof(s.Balance) /*+ unsafe.Sizeof(s.NLogs) + unsafe.Sizeof(s.NTokens) + unsafe.Sizeof(s.NErrors)*/ + unsafe.Sizeof(s.NItems) /*+ unsafe.Sizeof(s.NTotal)*/
	for _, record := range s.Items {
		size += unsafe.Sizeof(record)
	}
	return int(size)
}

// EXISTING_CODE
