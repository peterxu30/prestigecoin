package blockchain

import (
	"log"
	"testing"
)

const (
	Difficulty = 10 // Easy difficulty for testing purposes.
)

func TestBlockchain(t *testing.T) {
	log.Println("Test start.")

	bc, err := NewBlockChain(Difficulty)
	if err != nil {
		log.Println(err)
	}

	log.Println("Blockchain created.")

	msg1 := "John has 2 more PrestigeCoin than Jane"
	msg2 := "Jane has 10 more PrestigeCoin than David"

	err = bc.AddBlock(msg1)
	if err != nil {
		log.Println(err)
	}

	err = bc.AddBlock(msg2)
	if err != nil {
		log.Println(err)
	}

	log.Println("Blocks added.")

	bci := bc.Iterator()

	currBlock := bci.Next()
	currMsg := string(currBlock.GetData())
	if currMsg != msg2 {
		t.Errorf("Block held incorrect data. Expected: %s but got %s", msg2, currMsg)
	}

	currBlock = bci.Next()
	currMsg = string(currBlock.GetData())
	if string(currBlock.GetData()) != msg1 {
		t.Errorf("Block held incorrect data. Expected: %s but got %s", msg1, currMsg)
	}

	currBlock = bci.Next()
	currMsg = string(currBlock.GetData())
	if string(currBlock.GetData()) != "Genesis Block" {
		t.Errorf("Block held incorrect data. Expected: %s but got %s", "Genesis Block", currMsg)
	}
}
