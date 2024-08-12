// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package types

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type TransactionContainer struct {
	Address       base.Address            `json:"address"`
	Name          string                  `json:"name"`
	Balance       string                  `json:"balance"`
	NEvents       int64                   `json:"nEvents"`
	NTokens       int64                   `json:"nTokens"`
	NErrors       int64                   `json:"nErrors"`
	NTransactions int64                   `json:"nTransactions"`
	Items         []coreTypes.Transaction `json:"items"`
}

func (s *TransactionContainer) Summarize() {
	s.NTransactions = int64(len(s.Items))
	for _, tx := range s.Items {
		if tx.Receipt != nil {
			s.NEvents += int64(len(tx.Receipt.Logs))
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
		Address:       s.Address,
		Name:          s.Name,
		Balance:       s.Balance,
		NEvents:       s.NEvents,
		NTokens:       s.NTokens,
		NErrors:       s.NErrors,
		NTransactions: s.NTransactions,
	}
}
