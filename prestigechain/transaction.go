package prestigechain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const (
	newAchievement = "New Achievement"
)

type TXType int

const (
	Achievement TXType = iota
	Comparison
	Association
)

type Transaction struct {
	ID                     []byte
	User                   string
	Type                   TXType
	Value                  int
	Reason                 string
	RelevantTransactionIds [][]byte
}

func NewAchievementTX(user string, value int, reason string, relevantTxIds [][]byte) *Transaction {

	achievementReason := fmt.Sprintf("%d PrestigeCoins awarded to '%s' for %s", value, user, reason)

	tx := &Transaction{
		ID:     nil,
		User:   user,
		Type:   Achievement,
		Reason: achievementReason,
		Value:  value,
		RelevantTransactionIds: relevantTxIds,
	}
	tx.SetID()
	return tx
}

func (tx *Transaction) SetID() error {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		return err
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
	return nil
}

func AreEqualTransactionIds(id1 []byte, id2 []byte) bool {
	return bytes.Equal(id1, id2)
}

// func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
// 	return in.ScriptSig == unlockingData
// }

// func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
// 	return out.ScriptPubKey == unlockingData
// }

// func (tx Transaction) IsNewAchievement() bool {
// 	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
// }

func SerializeTXs(transactions []*Transaction) ([]byte, error) {
	var data bytes.Buffer
	encoder := gob.NewEncoder(&data)
	err := encoder.Encode(transactions)
	return data.Bytes(), err
}

func DeserializeTXs(encodedTransactions []byte) ([]*Transaction, error) {
	var transactions []*Transaction

	decoder := gob.NewDecoder(bytes.NewReader(encodedTransactions))
	err := decoder.Decode(&transactions)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
