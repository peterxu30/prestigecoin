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
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
	Type      TXType
	Reason    string
}

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func NewAchievementTX(value int, to, reason string) *Transaction {

	data := fmt.Sprintf("%d PrestigeCoins awarded to '%s' for %s", value, to, reason)

	txin := TXInput{[]byte{}, -1, newAchievement, Achievement, data}
	txout := TXOutput{value, to}
	tx := &Transaction{
		ID:   nil,
		Vin:  []TXInput{txin},
		Vout: []TXOutput{txout},
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
