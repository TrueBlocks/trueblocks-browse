package app

import (
	"fmt"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/sdk"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Block struct {
	BlockNumber  string   `json:"blockNumber"`
	Hash         string   `json:"hash"`
	Transactions []string `json:"transactions"`
}

func (a *App) GetBlock(bn uint64) Block {
	opts := sdk.BlocksOptions{
		BlockIds: []string{fmt.Sprintf("%d", bn)},
		Globals: sdk.Globals{
			Chain: "mainnet",
			Cache: true,
		},
	}
	blocks, _, err := opts.Blocks()
	if err != nil {
		runtime.EventsEmit(a.ctx, "error", err.Error())
		return Block{}
	}

	shrink := func(s string) string {
		return s[:6] + "..." + s[len(s)-4:]
	}

	ret := Block{
		BlockNumber: fmt.Sprintf("%d", blocks[0].BlockNumber),
		Hash:        shrink(blocks[0].Hash.Hex()),
	}

	line := []string{}
	for i := 0; i < len(blocks[0].Transactions); i++ {
		line = append(line, shrink(blocks[0].Transactions[i].Hash.Hex()))
		if (i+1)%8 == 0 {
			ret.Transactions = append(ret.Transactions, strings.Join(line, ", "))
			line = []string{}
		}
	}

	return ret
}
