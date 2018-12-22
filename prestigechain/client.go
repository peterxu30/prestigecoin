package prestigechain

type Client interface {
	AddNewAchievementTransaction(user, value int, reason string, relevantTxIds [][]byte) error
	FetchUpdates() error
	GetPrestigechain() *Prestigechain
}

// LocalClient contains the master Prestigechain. Not useful for more than testing.
type LocalClient struct {
	pc *Prestigechain
}

func NewLocalClient() *LocalClient {
	pc, err := NewPrestigechain()

	if err != nil {
		panic("Failed to create LocalClient.")
	}

	return &LocalClient{
		pc: pc,
	}
}

func (lc *LocalClient) AddNewAchievementTransaction(user string, reason string, value int, relevantTxIds [][]byte) error {
	tx := NewAchievementTX(user, value, reason, relevantTxIds)

	return lc.pc.AddBlock([]*Transaction{tx})
}

// Not implemented. LocalClient contains the master Prestigechain so no updates needed.
func (lc *LocalClient) FetchUpdates() error {
	return nil
}

func (lc *LocalClient) GetPrestigeChain() *Prestigechain {
	return lc.pc
}

func (lc *LocalClient) Delete() {
	lc.pc.Delete()
}
