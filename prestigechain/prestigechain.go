package prestigechain

import (
	blockchain "github.com/peterxu30/prestigecoin/blockchain"
)

const (
	difficulty             = 10
	genesisAddress         = "Genesis Block"
	genesisBlockCreatedMsg = "Mined Genesis Block"
)

type Prestigechain struct {
	bc *blockchain.Blockchain
}

type PrestigechainIterator struct {
	bci *blockchain.BlockchainIterator
}

func NewPrestigechain() (*Prestigechain, error) {
	genesisTx := NewAchievementTX(0, genesisAddress, genesisBlockCreatedMsg)
	encodedGenesisTx, err := SerializeTXs([]*Transaction{genesisTx})
	if err != nil {
		return nil, err
	}

	blockchain, err := blockchain.NewBlockChain(difficulty, encodedGenesisTx)
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
		block:        block,
		transactions: transactions,
	}, nil
}
