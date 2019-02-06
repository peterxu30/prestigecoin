package prestigechain

type Client interface {
	AddNewAchievementTransaction(username, value int, reason string, relevantTxIds [][]byte) error
	FetchUpdates() error
	GetPrestigechain() *Prestigechain
}

// MasterClient contains the master Prestigechain.
type MasterClient struct {
	pc *Prestigechain
	us *UserService
}

func NewMasterClient() (*MasterClient, error) {
	pc, err := NewPrestigechain()
	if err != nil {
		return nil, err
	}

	us, err := NewUserService()
	if err != nil {
		return nil, err
	}

	return &MasterClient{
		pc: pc,
		us: us,
	}, nil
}

func (mc *MasterClient) AddNewUser(username, password string) error {
	return mc.us.AddNewUser(username, password)
}

func (mc *MasterClient) UserExists(username string) bool {
	return mc.us.UserExists(username)
}

func (mc *MasterClient) ValidateUserPassword(username, password string) error {
	return mc.us.ValidateUserPassword(username, password)
}

// Assumes user has already been validated
func (mc *MasterClient) AddNewAchievementTransaction(username string, reason string, value int, relevantTxIds [][]byte) error {
	tx := NewAchievementTX(username, value, reason, relevantTxIds)

	return mc.pc.AddBlock([]*Transaction{tx})
}

// Not implemented. MasterClient contains the master Prestigechain so no updates needed.
func (mc *MasterClient) FetchUpdates() error {
	return nil
}

func (mc *MasterClient) GetPrestigeChain() *Prestigechain {
	return mc.pc
}

func (mc *MasterClient) Delete() {
	mc.pc.Delete()
}
