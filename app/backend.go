package app

import (
	"fmt"
	"strings"
)

type Block struct {
	BlockNumber  string   `json:"blockNumber"`
	Hash         string   `json:"hash"`
	Transactions []string `json:"transactions"`
}

func (a *App) GetBlock(bn uint64) Block {
	blk, _ := a.conn.GetBlockBodyByNumber(bn)

	ret := Block{
		Hash:        blk.Hash.Hex(),
		BlockNumber: fmt.Sprintf("%d", blk.BlockNumber),
	}
	cnt := 1
	four := []string{}
	for i := 0; i < len(blk.Transactions); i++ {
		four = append(four, shrink(blk.Transactions[i].Hash.Hex()))
		if cnt%8 == 0 {
			ret.Transactions = append(ret.Transactions, strings.Join(four, ", "))
			four = []string{}
		}
		cnt++
	}

	return ret
}

func shrink(s string) string {
	return s[:6] + "..." + s[len(s)-4:]
}

func (a *App) GetNames(page int) []string {
	first := page * 200
	last := first + 200
	n := a.namesArray[first:last]
	var ret []string
	for _, name := range n {
		ret = append(ret, fmt.Sprintf("%s: %s", name.Address.Hex(), name.Name))
	}
	return ret
}
