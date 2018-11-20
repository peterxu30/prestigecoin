package prestigechain

type Client interface {
	AddNewAchievementTransaction(user, value int, reason string, relevantTxIds [][]byte) error
	FetchUpdates() error
	GetPrestigechain() *Prestigechain
}

type BasicClient struct {
	pc *Prestigechain
}

func NewBasicClient() *BasicClient {
	pc, err := NewPrestigechain()

	if err != nil {
		panic("Failed to create BasicClient.")
	}

	return &BasicClient{
		pc: pc,
	}
}

func (bc *BasicClient) AddNewAchievementTransaction(user string, reason string, value int, relevantTxIds [][]byte) error {
	tx := NewAchievementTX(user, value, reason, relevantTxIds)

	return bc.pc.AddBlock([]*Transaction{tx})
}

// Not implemented
func (bc *BasicClient) FetchUpdates() error {
	return nil
}

func (bc *BasicClient) GetPrestigeChain() *Prestigechain {
	return bc.pc
}
