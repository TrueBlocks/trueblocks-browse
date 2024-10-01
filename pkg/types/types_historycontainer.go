package types

import (
	"encoding/json"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type HistoryContainer struct {
	Items      []coreTypes.Transaction `json:"items"`
	NItems     int                     `json:"nItems"`
	Address    base.Address            `json:"address"`
	Name       string                  `json:"name"`
	Balance    string                  `json:"balance"`
	NLogs      int                     `json:"nLogs"`
	NTokens    int                     `json:"nTokens"`
	NErrors    int                     `json:"nErrors"`
	Chain      string                  `json:"chain"`
	LastUpdate time.Time               `json:"lastUpdate"`
}

func NewHistoryContainer(chain string, address base.Address) HistoryContainer {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(chain), "monitors", address.Hex()+".mon.bin"))
	return HistoryContainer{
		Address:    address,
		Chain:      chain,
		LastUpdate: latest,
	}
}

func (s *HistoryContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *HistoryContainer) NeedsUpdate() bool {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(s.Chain), "monitors", s.Address.Hex()+".mon.bin"))
	if latest != s.LastUpdate {
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
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
	}
}

func (s *HistoryContainer) Summarize() {
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
}

func (s *HistoryContainer) SizeOf() int {
	size := unsafe.Sizeof(s.Address) + unsafe.Sizeof(s.Name) + unsafe.Sizeof(s.Balance) + unsafe.Sizeof(s.NLogs) + unsafe.Sizeof(s.NTokens) + unsafe.Sizeof(s.NErrors) + unsafe.Sizeof(s.NItems)
	for _, record := range s.Items {
		size += unsafe.Sizeof(record)
	}
	return int(size)
}
