package prestigechain

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/peterxu30/prestigecoin/utils"
	"github.com/stretchr/testify/assert"
)

const (
	user            = "tester"
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
	bytes := utils.GenerateRandomBytes(30)
	tx := NewAchievementTX(user, rand.Int(), utils.GenerateRandomString(maxStringLength), [][]byte{bytes})
	return tx
}

// Function is overkill. In general use, comparing if the transaction IDs will suffice for identity checks. (No reason to check for equality.)
func CompareTX(tx1 *Transaction, tx2 *Transaction) bool {
	if !bytes.Equal(tx1.ID, tx2.ID) {
		return false
	}

	if tx1.Type != tx2.Type {
		return false
	}

	if tx1.Reason != tx2.Reason {
		return false
	}

	if tx1.Value != tx2.Value {
		return false
	}

	tx1RelevantTransactionIds := tx1.RelevantTransactionIds
	tx2RelevantTransactionIds := tx2.RelevantTransactionIds

	if len(tx1RelevantTransactionIds) != len(tx2RelevantTransactionIds) {
		return false
	}

	for i, tx1reltxid := range tx1RelevantTransactionIds {
		tx2reltxid := tx2RelevantTransactionIds[i]

		if !bytes.Equal(tx1reltxid, tx2reltxid) {
			return false
		}
	}

	return true
}
