package blockchain

type Block interface {
	GetTimestamp() int64
	GetHash() []byte
	GetPreviousHash() []byte
	GetNonce() int
	GetDifficulty() int
	GetData() []byte
}

type BlockHeader interface {
	GetTimestamp() int64
	GetHash() []byte
	GetPreviousHash() []byte
	GetNonce() int
	GetDifficulty() int
}
