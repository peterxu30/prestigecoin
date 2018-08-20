package blockchain

type Blockchain struct {
	blocks     []Block
	difficulty int
}

func NewBlockChain(difficulty int) *Blockchain {
	blocks := []Block{NewGenesisSTBlock()}
	return &Blockchain{
		blocks:     blocks,
		difficulty: difficulty,
	}
}

func NewGenesisSTBlock() Block {
	return NewSTBlock(
		0,
		[]byte{},
		[]byte("Genesis Block"))
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewSTBlock(
		bc.difficulty,
		prevBlock.GetHash(),
		[]byte(data))
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) GetDifficulty() int {
	return bc.difficulty
}

//potentially dangerous, may need to return copies of the chain
func (bc *Blockchain) GetBlockchain() []Block {
	return bc.blocks
}
