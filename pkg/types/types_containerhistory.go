// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type HistoryContainer struct {
	Address    base.Address            `json:"address"`
	Balance    string                  `json:"balance"`
	NErrors    uint64                  `json:"nErrors"`
	NLogs      uint64                  `json:"nLogs"`
	NTokens    uint64                  `json:"nTokens"`
	NTotal     uint64                  `json:"nTotal"`
	Name       string                  `json:"name"`
	Items      []coreTypes.Transaction `json:"items"`
	NItems     uint64                  `json:"nItems"`
	Chain      string                  `json:"chain"`
	LastUpdate time.Time               `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewHistoryContainer(chain string, itemsIn []coreTypes.Transaction, address base.Address) HistoryContainer {
	ret := HistoryContainer{
		Items: make([]coreTypes.Transaction, 0, len(itemsIn)),
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getHistoryReload()
	// EXISTING_CODE
	ret.Address = address
	ret.LastUpdate, _ = ret.getHistoryReload() // it requires address
	// EXISTING_CODE
	return ret
}

func (s *HistoryContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *HistoryContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getHistoryReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *HistoryContainer) ShallowCopy() Containerer {
	return &HistoryContainer{
		Address:    s.Address,
		Balance:    s.Balance,
		NErrors:    s.NErrors,
		NLogs:      s.NLogs,
		NTokens:    s.NTokens,
		NTotal:     s.NTotal,
		Name:       s.Name,
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *HistoryContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	for _, tx := range s.Items {
		if tx.Receipt != nil {
			s.NLogs += uint64(len(tx.Receipt.Logs))
		}
		if tx.HasToken {
			s.NTokens++
		}
		if tx.IsError {
			s.NErrors++
		}
	}
	// EXISTING_CODE
}

func (s *HistoryContainer) getHistoryReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = file.MustGetLatestFileTime(filepath.Join(coreConfig.PathToCache(s.Chain), "monitors", s.Address.Hex()+".mon.bin"))
	reload = ret != s.LastUpdate
	// EXISTING_CODE
	return
}

type EveryTransactionFn func(item coreTypes.Transaction, data any) bool

func (s *HistoryContainer) ForEveryTransaction(process EveryTransactionFn, data any) bool {
	for _, item := range s.Items {
		if !process(item, data) {
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
