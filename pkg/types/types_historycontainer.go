package types

import (
	"unsafe"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type HistoryContainer struct {
	Items   []coreTypes.Transaction `json:"items"`
	NItems  int                     `json:"nItems"`
	Address base.Address            `json:"address"`
	Name    string                  `json:"name"`
	Balance string                  `json:"balance"`
	NLogs   int                     `json:"nLogs"`
	NTokens int                     `json:"nTokens"`
	NErrors int                     `json:"nErrors"`
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

func (s *HistoryContainer) ShallowCopy() HistoryContainer {
	return HistoryContainer{
		Address: s.Address,
		Name:    s.Name,
		Balance: s.Balance,
		NLogs:   s.NLogs,
		NTokens: s.NTokens,
		NErrors: s.NErrors,
		NItems:  s.NItems,
	}
}

func (s *HistoryContainer) SizeOf() int {
	size := unsafe.Sizeof(s.Address) + unsafe.Sizeof(s.Name) + unsafe.Sizeof(s.Balance) + unsafe.Sizeof(s.NLogs) + unsafe.Sizeof(s.NTokens) + unsafe.Sizeof(s.NErrors) + unsafe.Sizeof(s.NItems)
	for _, record := range s.Items {
		size += unsafe.Sizeof(record)
	}
	return int(size)
}
