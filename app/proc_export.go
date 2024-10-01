package app

import (
	"fmt"
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

func (a *App) ExportToCsv(addr string) {
	address, ok := a.ConvertToAddress(addr)
	if !ok {
		err := fmt.Errorf("Invalid address: " + addr)
		messages.SendError(a.ctx, err)
		return
	}

	isOpen := a.isFileOpen(address)
	if !isOpen {
		return
	}

	fn := fmt.Sprintf("history_%s.csv", address)
	lines := make([]string, 0, a.txCount(address)+2)

	// Add the CSV headers
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

	completed := a.forEveryTx(address, func(item types.Transaction) bool {
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
	})

	if !completed {
		messages.SendError(a.ctx, fmt.Errorf("export interrupted for address: %s", address.Hex()))
		return
	}

	file.StringToAsciiFile(fn, strings.Join(lines, "\n")+"\n")
	utils.System(fmt.Sprintf("open %s", fn))
}
