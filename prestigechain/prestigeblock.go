package prestigechain

import (
	blockchain "github.com/peterxu30/blockchain"
)

type PrestigeBlock struct {
	Header       *blockchain.BlockHeader
	Transactions []*Transaction
}

func NewPrestigeBlock(block *blockchain.Block) (*PrestigeBlock, error) {
	transactions, err := DeserializeTXs(block.GetData())
	if err != nil {
		return nil, err
	}

	return &PrestigeBlock{
		Header:       block.Header,
		Transactions: transactions,
	}, nil
}
