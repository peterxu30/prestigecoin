package blockchain

type Blockchain struct {
	blocks     []*Block
	difficulty int
}

func NewBlockChain(difficulty int) *Blockchain {
	blocks := []*Block{NewGenesisSTBlock()}
	return &Blockchain{
		blocks:     blocks,
		difficulty: difficulty,
	}
}

func NewGenesisSTBlock() *Block {
	return NewBlock(
		0,
		[]byte{},
		[]byte("Genesis Block"))
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(
		bc.difficulty,
		prevBlock.GetHash(),
		[]byte(data))
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) GetDifficulty() int {
	return bc.difficulty
}

// Potentially dangerous, may need to return copies of the chain. Counter: No one should have access to the source blockchain anyway.
func (bc *Blockchain) GetBlockchain() []*Block {
	return bc.blocks
}
