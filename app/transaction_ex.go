package app

// Find: NewViews
import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually, this will get put back into core

type TransactionEx struct {
	BlockNumber      base.Blknum    `json:"blockNumber"`
	TransactionIndex base.Txnum     `json:"transactionIndex"`
	Timestamp        base.Timestamp `json:"timestamp"`
	Date             string         `json:"date"`
	From             base.Address   `json:"from"`
	FromName         string         `json:"fromName"`
	To               base.Address   `json:"to"`
	ToName           string         `json:"toName"`
	Wei              base.Wei       `json:"wei"`
	Ether            string         `json:"ether"`
	Function         string         `json:"function"`
	HasToken         bool           `json:"hasToken"`
	IsError          bool           `json:"isError"`
	LogCount         uint64         `json:"logCount"`
}

func NewTransactionEx(a *App, tx *types.Transaction) *TransactionEx {
	fromName := a.namesMap[tx.From].Name.Name
	if len(fromName) == 0 {
		fromName = tx.From.String()
	} else if len(fromName) > 39 {
		fromName = fromName[:39] + "..."
	}
	toName := a.namesMap[tx.To].Name.Name
	if len(toName) == 0 {
		toName = tx.To.String()
	} else if len(toName) > 39 {
		toName = toName[:39] + "..."
	}
	ether := tx.Value.ToEtherStr(18)
	if tx.Value.IsZero() {
		ether = "-"
	} else if len(ether) > 5 {
		ether = ether[:5]
	}
	logCount := 0
	if tx.Receipt != nil {
		logCount = len(tx.Receipt.Logs)
	}

	return &TransactionEx{
		BlockNumber:      tx.BlockNumber,
		TransactionIndex: tx.TransactionIndex,
		Timestamp:        tx.Timestamp,
		Date:             tx.Date(),
		From:             tx.From,
		FromName:         fromName,
		To:               tx.To,
		ToName:           toName,
		Wei:              tx.Value,
		Ether:            ether,
		HasToken:         tx.HasToken,
		IsError:          tx.IsError,
		LogCount:         uint64(logCount),
		// Function:         tx.Function(),
	}
}
