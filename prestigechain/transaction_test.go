package prestigechain

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	to              = "tester"
	reason          = "testing"
	maxStringLength = 50
)

func TestTXEncodeAndDecodeHappyPath(t *testing.T) {
	tx := CreateTestTX()
	data, err := SerializeTXs([]*Transaction{tx})
	if err != nil {
		t.Errorf("Serializing failed with error %s", err.Error())
	}

	var txs []*Transaction
	txs, err = DeserializeTXs(data)
	if err != nil {
		t.Errorf("Deserializing failed with error %s", err.Error())
	}

	sameTx := txs[0]
	assert.True(t, CompareTX(tx, sameTx))
}

func CreateTestTX() *Transaction {
	tx := NewAchievementTX(rand.Int(), GenerateRandomString(), GenerateRandomString())
	return tx
}

func GenerateRandomString() string {
	length := rand.Intn(maxStringLength)
	var s string
	for i := 0; i < length; i++ {
		s += string('A' - 1 + rand.Intn(26))
	}
	return s
}

func CompareTX(tx1 *Transaction, tx2 *Transaction) bool {
	tx1ID := tx1.ID
	tx1Vin := tx1.Vin
	tx1Vout := tx1.Vout

	tx2ID := tx2.ID
	tx2Vin := tx2.Vin
	tx2Vout := tx2.Vout

	if !bytes.Equal(tx1ID, tx2ID) {
		return false
	}

	if len(tx1Vin) != len(tx2Vin) {
		return false
	}

	for i, txi := range tx1Vin {
		txi2 := tx2Vin[i]

		if !EqualTXInput(txi, txi2) {
			fmt.Println("F")
			return false
		}
	}

	if len(tx1Vout) != len(tx2Vout) {
		return false
	}

	for i, txo := range tx1Vout {
		txo2 := tx2Vout[i]

		if !EqualTXOutput(txo, txo2) {
			fmt.Println("G")
			return false
		}
	}

	return true
}

func EqualTXInput(in1 TXInput, in2 TXInput) bool {
	if in1.Reason != in2.Reason {
		return false
	}

	if in1.ScriptSig != in2.ScriptSig {
		return false
	}

	if !bytes.Equal(in1.Txid, in2.Txid) {
		return false
	}

	if in1.Type != in2.Type {
		return false
	}

	if in1.Vout != in2.Vout {
		return false
	}

	return true
}

func EqualTXOutput(out1 TXOutput, out2 TXOutput) bool {
	if out1.ScriptPubKey != out2.ScriptPubKey {
		return false
	}

	if out1.Value != out2.Value {
		return false
	}

	return true
}
