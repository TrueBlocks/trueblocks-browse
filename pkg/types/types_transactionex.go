package types

// Find: NewViews
import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// Find: NewViews

type TransactionEx struct {
	BlockNumber      base.Blknum  `json:"blockNumber"`
	Date             string       `json:"date"`
	Ether            string       `json:"ether"`
	From             base.Address `json:"from"`
	FromName         string       `json:"fromName"`
	Function         string       `json:"function"`
	HasToken         bool         `json:"hasToken"`
	IsError          bool         `json:"isError"`
	LogCount         uint64       `json:"logCount"`
	To               base.Address `json:"to"`
	ToName           string       `json:"toName"`
	TransactionIndex base.Txnum   `json:"transactionIndex"`
	Wei              base.Wei     `json:"wei"`
}

func NewTransactionEx(namesMap map[base.Address]NameEx, tx *coreTypes.Transaction) *TransactionEx {
	fromName := namesMap[tx.From].Name.Name
	if len(fromName) == 0 {
		fromName = tx.From.String()
	} else if len(fromName) > 39 {
		fromName = fromName[:39] + "..."
	}
	toName := namesMap[tx.To].Name.Name
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
		Date:             tx.Date(),
		Ether:            ether,
		From:             tx.From,
		FromName:         fromName,
		To:               tx.To,
		ToName:           toName,
		Wei:              tx.Value,
		HasToken:         tx.HasToken,
		IsError:          tx.IsError,
		LogCount:         uint64(logCount),
		// Function:         tx.Function(),
	}
}
