package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/sdk"
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
		},
	}
	blocks, _, _ := opts.Blocks()
	// blk, _ := a.conn.GetBlockBodyByNumber(bn)

	ret := Block{
		BlockNumber: fmt.Sprintf("%d", blocks[0].BlockNumber),
		Hash:        blocks[0].Hash.Hex(),
	}
	// 	Hash:        blk.Hash.Hex(),
	// 	BlockNumber: fmt.Sprintf("%d", blk.BlockNumber),
	// }
	// cnt := 1
	// four := []string{}
	// for i := 0; i < len(blk.Transactions); i++ {
	// 	four = append(four, shrink(blk.Transactions[i].Hash.Hex()))
	// 	if cnt%8 == 0 {
	// 		ret.Transactions = append(ret.Transactions, strings.Join(four, ", "))
	// 		four = []string{}
	// 	}
	// 	cnt++
	// }

	return ret
}

// func shrink(s string) string {
// 	return s[:6] + "..." + s[len(s)-4:]
// }

func (a *App) GetNames(page int) []string {
	first := page
	last := first + 20
	if len(a.namesArray) < last {
		return []string{"No names loaded"}
	}
	n := a.namesArray[first:last]
	var ret []string
	for _, name := range n {
		ret = append(ret, fmt.Sprintf("%s: %s", name.Address.Hex(), name.Name))
	}
	return ret
}
