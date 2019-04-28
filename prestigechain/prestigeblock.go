package prestigechain

import (
	cloudchain "github.com/peterxu30/cloudchain"
)

type PrestigeBlock struct {
	Header       *cloudchain.BlockHeader
	Transactions []*Transaction
}

func NewPrestigeBlock(block *cloudchain.Block) (*PrestigeBlock, error) {
	transactions, err := DeserializeTXs(block.Data)
	if err != nil {
		return nil, err
	}

	return &PrestigeBlock{
		Header:       block.Header,
		Transactions: transactions,
	}, nil
}
