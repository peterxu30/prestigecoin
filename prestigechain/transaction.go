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

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
	Reason    string
}

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func NewAchievementTX(value int, to, reason string) *Transaction {

	data := fmt.Sprintf("%d PrestigeCoins awarded to '%s' for %s", value, to, reason)

	txin := TXInput{[]byte{}, -1, newAchievement, data}
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
