package prestigechain

import (
	blockchain "github.com/peterxu30/prestigecoin/blockchain"
)

type PrestigeBlock struct {
	block        *blockchain.Block
	transactions []*Transaction
}
