package prestigechain

import (
	bc "github.com/peterxu30/prestigecoin/blockchain"
)

const (
	difficulty = 10
)

type Prestigechain struct {
	bc *bc.Blockchain
}

func NewPrestigechain() *Prestigechain {
	blockchain, _ := bc.NewBlockChain(difficulty)
	return &Prestigechain{blockchain}
}
