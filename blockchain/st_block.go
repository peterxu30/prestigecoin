package blockchain

import (
	"time"
)

//Single Transaction Block
type STBlock struct {
	header *STBlockHeader
	data   []byte
}

type STBlockHeader struct {
	timestamp  int64
	hash       []byte
	prevHash   []byte
	nonce      int
	difficulty int
}

func NewSTBlock(difficulty int, prevHash []byte, data []byte) *STBlock {
	header := &STBlockHeader{
		timestamp:  time.Now().Unix(),
		prevHash:   prevHash,
		difficulty: difficulty,
	}

	stb := &STBlock{
		header: header,
		data:   data,
	}

	pow := NewProofOfWork(stb, difficulty)
	nonce, hash := pow.Run()

	header.hash = hash
	header.nonce = nonce

	return stb
}

func (stb *STBlock) GetTimestamp() int64 {
	return stb.header.GetTimestamp()
}

func (stb *STBlock) GetHash() []byte {
	return stb.header.GetHash()
}

func (stb *STBlock) GetPreviousHash() []byte {
	return stb.header.GetPreviousHash()
}

func (stb *STBlock) GetNonce() int {
	return stb.header.GetNonce()
}

func (stb *STBlock) GetDifficulty() int {
	return stb.header.GetDifficulty()
}

func (stb *STBlock) GetData() []byte {
	return stb.data
}

func (stbh *STBlockHeader) GetTimestamp() int64 {
	return stbh.timestamp
}

func (stbh *STBlockHeader) GetHash() []byte {
	return stbh.hash
}

func (stbh *STBlockHeader) GetPreviousHash() []byte {
	return stbh.prevHash
}

func (stbh *STBlockHeader) GetNonce() int {
	return stbh.nonce
}

func (stbh *STBlockHeader) GetDifficulty() int {
	return stbh.difficulty
}
