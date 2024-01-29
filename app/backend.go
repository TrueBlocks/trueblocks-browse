package app

import (
	"fmt"
)

type Block struct {
	BlockNumber string `json:"blockNumber"`
	Hash        string `json:"hash"`
}

func (a *App) GetBlock(bn uint64) Block {
	blk, _ := a.Conn.GetBlockBodyByNumber(bn)
	return Block{
		Hash:        blk.Hash.Hex(),
		BlockNumber: fmt.Sprintf("%d", blk.BlockNumber),
	}
}
