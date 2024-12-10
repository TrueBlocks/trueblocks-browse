// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type IncomingContainer struct {
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

func NewIncomingContainer(chain string, itemsIn []Transaction, address base.Address) IncomingContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := IncomingContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewIncomingUpdater(chain, address),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewIncomingUpdater(chain string, address base.Address, resetIn ...bool) sdk.Updater {
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
	u, _ := sdk.NewUpdater("incoming", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *IncomingContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *IncomingContainer) GetItems() interface{} {
	return s.Items
}

func (s *IncomingContainer) SetItems(items interface{}) {
	s.Items = items.([]Transaction)
}

func (s *IncomingContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *IncomingContainer) ShallowCopy() Containerer {
	ret := &IncomingContainer{
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

func (s *IncomingContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *IncomingContainer) passesFilter(item *Transaction, filter *Filter) (ret bool) {
	_ = item // linter
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *IncomingContainer) Accumulate(item *Transaction) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *IncomingContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *IncomingContainer) CollateAndFilter(filter *Filter) interface{} {
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

func (s *IncomingContainer) ForEveryItem(process EveryTransactionFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *IncomingContainer) Sort() (err error) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
