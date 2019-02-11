package client

import (
	"log"
	"sync"

	"github.com/peterxu30/prestigecoin/prestigechain"
)

type Client interface {
	AddNewAchievementTransaction(username, value int, reason string, relevantTxIds [][]byte) error
	FetchUpdates() error
	GetPrestigechain() *prestigechain.Prestigechain
}

// TODO: Make the MasterClient thread safe

// MasterClient contains the master Prestigechain.
type MasterClient struct {
	pc *prestigechain.Prestigechain
	us *UserService
}

var _masterClient *MasterClient
var once sync.Once

func GetOrCreateMasterClient() *MasterClient {
	once.Do(func() {
		var err error
		_masterClient, err = newMasterClient()
		if err != nil {
			log.Panicln("Failed to create MasterClient", err)
		}
	})

	return _masterClient
}

func newMasterClient() (*MasterClient, error) {
	pc, err := prestigechain.NewPrestigechain()
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
	tx := prestigechain.NewAchievementTX(username, value, reason, relevantTxIds)
	return mc.pc.AddBlock([]*prestigechain.Transaction{tx})
}

// Not implemented. MasterClient contains the master Prestigechain so no updates needed.
func (mc *MasterClient) FetchUpdates() error {
	return nil
}

func (mc *MasterClient) GetPrestigeChain() *prestigechain.Prestigechain {
	return mc.pc
}

func DeleteMasterClient(mc *MasterClient) {
	prestigechain.DeletePrestigechain(mc.pc)
	DeleteUserService(mc.us)
}
