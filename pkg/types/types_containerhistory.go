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
	NTotal     int                     `json:"nTotal"`
	Name       string                  `json:"name"`
	Balance    string                  `json:"balance"`
	NLogs      int                     `json:"nLogs"`
	NTokens    int                     `json:"nTokens"`
	NErrors    int                     `json:"nErrors"`
	Items      []coreTypes.Transaction `json:"items"`
	NItems     int                     `json:"nItems"`
	Chain      string                  `json:"chain"`
	LastUpdate time.Time               `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewHistoryContainer(chain string, address base.Address) HistoryContainer {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(chain), "monitors", address.Hex()+".mon.bin"))
	ret := HistoryContainer{
		Chain:      chain,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Address = address
	// EXISTING_CODE
	return ret
}

func (s *HistoryContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *HistoryContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(s.Chain), "monitors", s.Address.Hex()+".mon.bin"))
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *HistoryContainer) ShallowCopy() Containerer {
	return &HistoryContainer{
		Address:    s.Address,
		Name:       s.Name,
		Balance:    s.Balance,
		NLogs:      s.NLogs,
		NTokens:    s.NTokens,
		NErrors:    s.NErrors,
		NItems:     s.NItems,
		NTotal:     s.NTotal,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *HistoryContainer) Summarize() {
	// EXISTING_CODE
	s.NItems = len(s.Items)
	for _, tx := range s.Items {
		if tx.Receipt != nil {
			s.NLogs += len(tx.Receipt.Logs)
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

func HistoryX() {
	// EXISTING_CODE
	// EXISTING_CODE
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
