package blockchain

import (
	"time"
)

type Block struct {
	header *BlockHeader
	data   []byte
}

type BlockHeader struct {
	timestamp  int64
	hash       []byte
	prevHash   []byte
	nonce      int
	difficulty int
}

func NewBlock(difficulty int, prevHash []byte, data []byte) *Block {
	header := &BlockHeader{
		timestamp:  time.Now().Unix(),
		prevHash:   prevHash,
		difficulty: difficulty,
	}

	block := &Block{
		header: header,
		data:   data,
	}

	pow := NewProofOfWork(block, difficulty)
	nonce, hash := pow.Run()

	header.hash = hash
	header.nonce = nonce

	return block
}

func (block *Block) GetTimestamp() int64 {
	return block.header.GetTimestamp()
}

func (block *Block) GetHash() []byte {
	return block.header.GetHash()
}

func (block *Block) GetPreviousHash() []byte {
	return block.header.GetPreviousHash()
}

func (block *Block) GetNonce() int {
	return block.header.GetNonce()
}

func (block *Block) GetDifficulty() int {
	return block.header.GetDifficulty()
}

func (block *Block) GetData() []byte {
	return block.data
}

func (header *BlockHeader) GetTimestamp() int64 {
	return header.timestamp
}

func (header *BlockHeader) GetHash() []byte {
	return header.hash
}

func (header *BlockHeader) GetPreviousHash() []byte {
	return header.prevHash
}

func (header *BlockHeader) GetNonce() int {
	return header.nonce
}

func (header *BlockHeader) GetDifficulty() int {
	return header.difficulty
}
