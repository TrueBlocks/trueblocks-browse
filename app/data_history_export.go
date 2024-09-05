package app

import (
	"fmt"
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

func (a *App) ExportToCsv(addr string) {
	address, ok := a.ConvertToAddress(addr)
	if !ok {
		err := fmt.Errorf("Invalid address: " + addr)
		messages.SendError(a.ctx, err)
		return
	}

	historyMutex.Lock()
	_, exists := a.historyMap[address]
	historyMutex.Unlock()

	if exists {
		fn := fmt.Sprintf("history_%s.csv", address)
		lines := make([]string, 0, len(a.historyMap[address].Items)+2)
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
		historyMutex.Lock()
		for _, item := range a.historyMap[address].Items {
			lines = append(lines, fmt.Sprintf("%d,%s,%d,%s,%s,%s,%d,%d,%d,%d,%d,%d,%s,%s",
				item.BlockNumber,
				item.BlockHash.Hex(),
				item.TransactionIndex,
				item.Hash.Hex(),
				item.From.Hex(),
				item.To.Hex(),
				item.Value.Uint64(),
				item.Gas,
				item.GasPrice,
				item.GasUsed,
				item.Timestamp,
				item.Nonce,
				item.Input,
				item.TransactionType))
		}
		historyMutex.Unlock()
		file.StringToAsciiFile(fn, strings.Join(lines, "\n")+"\n")
		utils.System(fmt.Sprintf("open %s", fn))
	}
}
