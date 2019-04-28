package prestigechain

import (
	"context"

	cloudchain "github.com/peterxu30/cloudchain"
)

const (
	prestigeChainDbDir     = ".pcdb"
	difficulty             = 10
	genesisBlockCreator    = "The Architects"
	genesisBlockCreatedMsg = "Mined Genesis Block"
)

type Prestigechain struct {
	cc *cloudchain.CloudChain
}

type PrestigechainIterator struct {
	cci *cloudchain.CloudChainIterator
}

func NewPrestigechain(ctx context.Context, projectId string) (*Prestigechain, error) {
	genesisTx := NewAchievementTX(genesisBlockCreator, 0, genesisBlockCreatedMsg, nil)
	encodedGenesisTx, err := SerializeTXs([]*Transaction{genesisTx})
	if err != nil {
		return nil, err
	}

	cloudchain, err := cloudchain.NewCloudChain(ctx, projectId, difficulty, encodedGenesisTx)
	if err != nil {
		return nil, err
	}

	return &Prestigechain{cloudchain}, nil
}

func (pc *Prestigechain) AddBlock(ctx context.Context, transactions []*Transaction) (*PrestigeBlock, error) {
	encodedTransactions, err := SerializeTXs(transactions)
	if err != nil {
		return nil, err
	}

	block, err := pc.cc.AddBlock(ctx, encodedTransactions)
	if err != nil {
		return nil, err
	}

	return NewPrestigeBlock(block)
}

func (pc *Prestigechain) Difficulty() int {
	return pc.cc.Difficulty()
}

// func (pc *Prestigechain) Close() error {
// 	return pc.cc.Close()
// }

func (pc *Prestigechain) Iterator(ctx context.Context) (*PrestigechainIterator, error) {
	cci, err := pc.cc.Iterator(ctx)
	if err != nil {
		return nil, err
	}

	pci := &PrestigechainIterator{cci}
	return pci, nil
}

// Consider deleting
func (pc *Prestigechain) GetTransactionById(ctx context.Context, id []byte) (*Transaction, error) {
	pci, err := pc.Iterator(ctx)
	if err != nil {
		return nil, err
	}

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

func DeletePrestigechain(ctx context.Context, pc *Prestigechain) {
	cloudchain.DeleteCloudChain(ctx, pc.cc)
}

func (pci *PrestigechainIterator) Next() (*PrestigeBlock, error) {
	block, err := pci.cci.Next()
	if block == nil {
		return nil, err
	}

	return NewPrestigeBlock(block)
}
