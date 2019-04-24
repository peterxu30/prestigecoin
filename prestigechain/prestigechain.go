package prestigechain

import (
	blockchain "github.com/peterxu30/blockchain"
)

const (
	prestigeChainDbDir     = ".pcdb"
	difficulty             = 10
	genesisBlockCreator    = "The Architects"
	genesisBlockCreatedMsg = "Mined Genesis Block"
)

type Prestigechain struct {
	bc *blockchain.Blockchain
}

type PrestigechainIterator struct {
	bci *blockchain.BlockchainIterator
}

func NewPrestigechain() (*Prestigechain, error) {
	genesisTx := NewAchievementTX(genesisBlockCreator, 0, genesisBlockCreatedMsg, nil)
	encodedGenesisTx, err := SerializeTXs([]*Transaction{genesisTx})
	if err != nil {
		return nil, err
	}

	blockchain, err := blockchain.NewBlockChain(prestigeChainDbDir, difficulty, encodedGenesisTx)
	if err != nil {
		return nil, err
	}

	return &Prestigechain{blockchain}, nil
}

func (pc *Prestigechain) AddBlock(transactions []*Transaction) (*PrestigeBlock, error) {
	encodedTransactions, err := SerializeTXs(transactions)
	if err != nil {
		return nil, err
	}

	block, err := pc.bc.AddBlock(encodedTransactions)
	if err != nil {
		return nil, err
	}

	return NewPrestigeBlock(block)
}

func (pc *Prestigechain) Difficulty() int {
	return pc.bc.Difficulty()
}

func (pc *Prestigechain) Close() error {
	return pc.bc.Close()
}

func (pc *Prestigechain) Iterator() *PrestigechainIterator {
	pci := &PrestigechainIterator{pc.bc.Iterator()}
	return pci
}

// Consider deleting
func (pc *Prestigechain) GetTransactionById(id []byte) (*Transaction, error) {
	pci := pc.Iterator()

	for {
		block, err := pci.Next()
		if err != nil {
			return nil, err
		}

		transactions := block.Transactions
		for _, transaction := range transactions {
			if AreEqualTransactionIds(transaction.ID, id) {
				return transaction, nil
			}
		}

		if block.Header.IsLastBlock() {
			break
		}
	}

	return nil, nil
}

func DeletePrestigechain(pc *Prestigechain) {
	blockchain.DeleteBlockchain(pc.bc)
}

func (pci *PrestigechainIterator) Next() (*PrestigeBlock, error) {
	block, err := pci.bci.Next()
	if block == nil {
		return nil, err
	}

	return NewPrestigeBlock(block)
}
