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

func (pc *Prestigechain) AddBlock(transactions []*Transaction) error {
	encodedTransactions, err := SerializeTXs(transactions)
	if err != nil {
		return err
	}

	return pc.bc.AddBlock(encodedTransactions)
}

func (pc *Prestigechain) GetDifficulty() int {
	return pc.bc.GetDifficulty()
}

func (pc *Prestigechain) Close() error {
	return pc.bc.Close()
}

func (pc *Prestigechain) Iterator() *PrestigechainIterator {
	pci := &PrestigechainIterator{pc.bc.Iterator()}
	return pci
}

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

		if block.Block.IsLastBlock() {
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
	if err != nil {
		return nil, err
	}

	transactions, err := DeserializeTXs(block.GetData())
	if err != nil {
		return nil, err
	}

	return &PrestigeBlock{
		Block:        block,
		Transactions: transactions,
	}, nil
}
