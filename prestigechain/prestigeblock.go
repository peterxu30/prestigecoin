package prestigechain

import (
	blockchain "github.com/peterxu30/blockchain"
)

type PrestigeBlock struct {
	Block        *blockchain.Block
	Transactions []*Transaction
}
