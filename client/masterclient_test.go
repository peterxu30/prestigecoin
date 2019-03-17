package client

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/peterxu30/prestigecoin/prestigechain"
	"github.com/peterxu30/prestigecoin/utils"
)

const (
	clientUser = "clientUser"
)

func TestCreateAndUseMasterClientHappyPath(t *testing.T) {
	client := GetOrCreateMasterClient()
	defer DeleteMasterClient(client)

	genBytes := utils.GenerateRandomBytes(30)
	testRelTxIds := [][]byte{genBytes}
	value := rand.Int()
	reason := "happy testing"
	updateData := PrestigechainUpdateData{
		User:   clientUser,
		Reason: reason,
		Value:  value,
		RelevantTransactionIds: testRelTxIds,
	}
	err := client.AddNewAchievementTransaction(updateData)
	if err != nil {
		t.Errorf("Failed to add new achievement transaction")
	}

	pc := client.GetPrestigeChain()
	pci := pc.Iterator()

	head, err := pci.Next()
	if err != nil {
		t.Errorf("Failed to add retrieve new block")
	}

	if len(head.Transactions) != 1 {
		t.Errorf("Incorrect number of transactions. Expected 1, got %v", len(head.Transactions))
	}

	tx := head.Transactions[0]
	if len(tx.RelevantTransactionIds) != 1 {
		t.Errorf("Incorrect number of transaction Ids. Expected 1, got %v", len(tx.RelevantTransactionIds))
	}

	if !bytes.Equal(tx.RelevantTransactionIds[0], genBytes) {
		t.Errorf("Incorrect relevant transaction id.")
	}

	if tx.Type != prestigechain.Achievement {
		t.Errorf("Incorrect transaction type.")
	}

	if tx.User != clientUser {
		t.Errorf("Incorrect user.")
	}

	if tx.Value != value {
		t.Errorf("Incorrect value.")
	}

	fullReason := fmt.Sprintf("%v PrestigeCoins awarded to '%s' for %s", value, clientUser, reason)
	if tx.Reason != fullReason {
		t.Errorf("Incorrect reason.")
	}
}
