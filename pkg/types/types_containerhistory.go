package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"sync"
	"time"
	"unsafe"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
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
	ret = utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(s.Chain), "monitors", s.Address.Hex()+".mon.bin"))
	reload = ret != s.LastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (s *HistoryContainer) SizeOf() int {
	size := unsafe.Sizeof(s.Address) + unsafe.Sizeof(s.Name) + unsafe.Sizeof(s.Balance) + unsafe.Sizeof(s.NLogs) + unsafe.Sizeof(s.NTokens) + unsafe.Sizeof(s.NErrors) + unsafe.Sizeof(s.NItems) + unsafe.Sizeof(s.NTotal)
	for _, record := range s.Items {
		size += unsafe.Sizeof(record)
	}
	return int(size)
}

type HistoryMap struct {
	internal sync.Map
}

func (h *HistoryMap) Store(address base.Address, historyContainer HistoryContainer) {
	h.internal.Store(address, historyContainer)
}

func (h *HistoryMap) Load(address base.Address) (HistoryContainer, bool) {
	value, ok := h.internal.Load(address)
	if !ok {
		return HistoryContainer{}, false
	}
	return value.(HistoryContainer), true
}

func (h *HistoryMap) Delete(address base.Address) {
	h.internal.Delete(address)
}

func (h *HistoryMap) Range(f func(address base.Address, historyContainer HistoryContainer) bool) {
	h.internal.Range(func(key, value interface{}) bool {
		return f(key.(base.Address), value.(HistoryContainer))
	})
}

// EXISTING_CODE
