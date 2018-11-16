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

// // Finds transactions with unspent Prestigecoins of the given address.
// func (pc *Prestigechain) FindUnspentTransactions(address string) ([]Transaction, error) {
// 	var unspentTXs []Transaction
// 	spentTXOs := make(map[string][]int)
// 	pci := pc.Iterator()

// 	for {
// 		block, err := pci.Next()
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, tx := range block.Transactions {
// 			txID := hex.EncodeToString(tx.ID)

// 		Outputs:
// 			for outIdx, out := range tx.Vout {
// 				// Was the output spent?
// 				if spentTXOs[txID] != nil {
// 					for _, spentOut := range spentTXOs[txID] {
// 						if spentOut == outIdx {
// 							continue Outputs
// 						}
// 					}
// 				}

// 				if out.CanBeUnlockedWith(address) {
// 					unspentTXs = append(unspentTXs, *tx)
// 				}
// 			}

// 			if tx.IsNewAchievement() == false {
// 				for _, in := range tx.Vin {
// 					if in.CanUnlockOutputWith(address) {
// 						inTxID := hex.EncodeToString(in.Txid)
// 						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
// 					}
// 				}
// 			}
// 		}

// 		if len(block.Block.GetPreviousHash()) == 0 {
// 			break
// 		}
// 	}
// 	return unspentTXs, nil
// }

// func (pc *Prestigechain) FindUnspentTransactionOutputs(address string) ([]TXOutput, error) {
// 	var UTXOs []TXOutput
// 	unspentTransactions, err := pc.FindUnspentTransactions(address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, tx := range unspentTransactions {
// 		for _, out := range tx.Vout {
// 			if out.CanBeUnlockedWith(address) {
// 				UTXOs = append(UTXOs, out)
// 			}
// 		}
// 	}

// 	return UTXOs, nil
// }

// func NewUTXOTransaction(from, to string, amount int, reason string, pc *Prestigechain) (*Transaction, error) {
// 	var inputs []TXInput
// 	var outputs []TXOutput

// 	acc, validOutputs, err := pc.FindSpendableOutputs(from, amount)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if acc < amount {
// 		return nil, errors.New("Insufficient funds to create transaction.")
// 	}

// 	for txid, outs := range validOutputs {
// 		txID, err := hex.DecodeString(txid)

// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, out := range outs {
// 			input := TXInput{txID, out, from, Achievement, reason}
// 			inputs = append(inputs, input)
// 		}
// 	}

// 	outputs = append(outputs, TXOutput{amount, to})
// 	if acc > amount {
// 		outputs = append(outputs, TXOutput{acc - amount, from})
// 	}

// 	tx := Transaction{nil, inputs, outputs}
// 	tx.SetID()

// 	return &tx, nil
// }

// func (pc *Prestigechain) FindSpendableOutputs(address string, amount int) (int, map[string][]int, error) {
// 	unspentOutputs := make(map[string][]int)
// 	unspentTXs, err := pc.FindUnspentTransactions(address)
// 	if err != nil {
// 		return 0, nil, err
// 	}

// 	accumulated := 0

// Work:
// 	for _, tx := range unspentTXs {
// 		txID := hex.EncodeToString(tx.ID)

// 		for outIdx, out := range tx.Vout {
// 			if out.CanBeUnlockedWith(address) && accumulated < amount {
// 				accumulated += out.Value
// 				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

// 				if accumulated >= amount {
// 					break Work
// 				}
// 			}
// 		}
// 	}

// 	return accumulated, unspentOutputs, nil
// }

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
