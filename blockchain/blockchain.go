package blockchain

import (
	"os"

	"github.com/boltdb/bolt"
)

const (
	dbDir        = ".db"
	dbFile       = "main.db"
	blocksBucket = "blocksBucket"
	headBlock    = "head"
)

type Blockchain struct {
	head       []byte
	difficulty int
	db         *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func NewBlockChain(difficulty int) (*Blockchain, error) {
	return newBlockChainWithDbPath(difficulty, dbDir)
}

func newBlockChainWithDbPath(difficulty int, dbPath string) (*Blockchain, error) {
	fullDbPath := dbPath + "/" + dbFile
	if _, err := os.Stat(fullDbPath); os.IsNotExist(err) {
		err = os.MkdirAll(dbPath, 0700)
		if err != nil {
			return nil, err
		}
	}

	db, err := bolt.Open(fullDbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	var encodedHead []byte
	err = db.Update(func(tx *bolt.Tx) error {
		blocksBucketKey := []byte(blocksBucket)
		headBlockKey := []byte(headBlock)

		b := tx.Bucket(blocksBucketKey)

		if b == nil {
			b, err = tx.CreateBucket(blocksBucketKey)
			if err != nil {
				return err
			}

			genesisBlock := NewGenesisBlock()
			encodedGenesisBlock, err := genesisBlock.Serialize()

			if err != nil {
				return err
			}

			err = b.Put(genesisBlock.GetHash(), encodedGenesisBlock)
			if err != nil {
				return err
			}

			err = b.Put(headBlockKey, genesisBlock.GetHash())
			if err != nil {
				return err
			}

			encodedHead = genesisBlock.GetHash()
		} else {
			encodedHead = b.Get(headBlockKey)
		}

		return nil
	})

	blockchain := &Blockchain{
		head:       encodedHead,
		difficulty: difficulty,
		db:         db,
	}

	return blockchain, nil
}

func NewGenesisBlock() *Block {
	return NewBlock(
		0,
		nil,
		[]byte("Genesis Block"))
}

func (bc *Blockchain) AddBlock(data []byte) error {
	newBlock := NewBlock(bc.difficulty, bc.head, data)

	err := bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedNewBlock, err := newBlock.Serialize()

		if err != nil {
			return err
		}

		err = b.Put(newBlock.GetHash(), encodedNewBlock)

		if err != nil {
			return err
		}

		err = b.Put([]byte(headBlock), newBlock.GetHash())
		bc.head = newBlock.GetHash()
		return nil
	})

	return err
}

func (bc *Blockchain) GetDifficulty() int {
	return bc.difficulty
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.head, bc.db}
	return bci
}

func (bci *BlockchainIterator) Next() *Block {
	if bci.currentHash == nil {
		return nil
	}

	var encodedBlock []byte

	err := bci.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock = b.Get(bci.currentHash)

		return nil
	})

	block, err := DeserializeBlock(encodedBlock)

	if err != nil {
		panic(err) // consider switching to return error
	}

	bci.currentHash = block.GetPreviousHash()

	return block
}
