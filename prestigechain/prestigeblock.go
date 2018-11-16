package prestigechain

import (
	blockchain "github.com/peterxu30/prestigecoin/blockchain"
)

type PrestigeBlock struct {
	Block        *blockchain.Block
	Transactions []*Transaction
}
