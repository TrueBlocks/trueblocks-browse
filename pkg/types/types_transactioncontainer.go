package types

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type TransactionContainer struct {
	Items   []coreTypes.Transaction `json:"items"`
	NItems  int                     `json:"nItems"`
	Address base.Address            `json:"address"`
	Name    string                  `json:"name"`
	Balance string                  `json:"balance"`
	NEvents int                     `json:"nEvents"`
	NTokens int                     `json:"nTokens"`
	NErrors int                     `json:"nErrors"`
}

func (s *TransactionContainer) Summarize() {
	s.NItems = len(s.Items)
	for _, tx := range s.Items {
		if tx.Receipt != nil {
			s.NEvents += len(tx.Receipt.Logs)
		}
		if tx.HasToken {
			s.NTokens++
		}
		if tx.IsError {
			s.NErrors++
		}
	}
}

func (s *TransactionContainer) ShallowCopy() TransactionContainer {
	return TransactionContainer{
		Address: s.Address,
		Name:    s.Name,
		Balance: s.Balance,
		NEvents: s.NEvents,
		NTokens: s.NTokens,
		NErrors: s.NErrors,
		NItems:  s.NItems,
	}
}
