package blockchain

import (
	"fmt"
	"testing"
)

const (
	Difficulty = 24
)

func TestBlockchain(t *testing.T) {

	bc := NewBlockChain(Difficulty)

	msg1 := "John has 2 more PrestigeCoin than Jane"
	msg2 := "Jane has 10 more PrestigeCoin than David"

	bc.AddBlock("John has 2 more PrestigeCoin than Jane")
	bc.AddBlock("Jane has 10 more PrestigeCoin than David")

	blocks := bc.GetBlockchain()

	if len(blocks) != 3 {
		t.Error("Expected 3 blocks, got ", len(blocks))
	}

	for _, block := range blocks {
		fmt.Printf("Data: %s\n", block.GetData())
	}

	data1 := string(blocks[1].GetData())
	if data1 != msg1 {
		t.Error("Block1 held incorrect data")
	}

	data2 := string(blocks[2].GetData())
	if data2 != msg2 {
		t.Error("Block2 held incorrect data")
	}

}
