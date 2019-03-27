package prestigechain

import (
	blockchain "github.com/peterxu30/blockchain"
)

type PrestigeBlock struct {
	Header       *blockchain.BlockHeader
	Transactions []*Transaction
}
