package app

import (
	"fmt"
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

// TODO: If this uses chifra export, it could be much expanded to include
// TODO: all the options that chifra export has. Also, this should stream
// TODO: it output to a Writer not a string. Much fast for much larger files.
func (a *App) ExportAddress(address base.Address) {
	isOpen := a.isFileOpen(address)
	if !isOpen {
		return
	}

	fn := fmt.Sprintf("history_%s.csv", address)
	lines := make([]string, 0, a.txCount(address)+2)

	exportLine := func(item coreTypes.Transaction, data any) bool {
		if len(lines) == 0 {
			lines = append(lines, fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
				"BlockNumber",
				"BlockHash",
				"TransactionIndex",
				"Hash",
				"From",
				"To",
				"Value.Uint64()",
				"Gas",
				"GasPrice",
				"GasUsed",
				"Timestamp",
				"Nonce",
				"Input",
				"TransactionType"))
		}
		lines = append(lines, fmt.Sprintf("%d,%d,%s,%d,%s,%s,%d,%d,%d,%d,%s,%s,%d,%s,%s",
			item.BlockNumber,
			item.TransactionIndex,
			item.Date(),
			item.Timestamp,
			item.From.Hex(),
			item.To.Hex(),
			item.Value.Uint64(),
			item.Gas,
			item.GasPrice,
			item.GasUsed,
			item.BlockHash.Hex(),
			item.Hash.Hex(),
			item.Nonce,
			utils.FormattedCode(false, item.Input),
			item.TransactionType))
		return true
	}

	h, _ := a.HistoryCache.Load(address)
	completed := h.ForEveryTransaction(exportLine, nil)
	if !completed {
		err := fmt.Errorf("export for %s interrupted after %d lines", address.Hex(), len(lines))
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: err.Error(),
		})
		return
	}

	file.StringToAsciiFile(fn, strings.Join(lines, "\n")+"\n")
	utils.System(fmt.Sprintf("open %s", fn))
}
