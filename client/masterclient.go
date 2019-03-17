package client

import (
	"log"
	"sync"

	"github.com/peterxu30/prestigecoin/prestigechain"
)

// TODO: Make the MasterClient thread safe. Might already be since BoltDB is threadsafe

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
		log.Println("bad pc")
		return nil, err
	}

	us, err := NewUserService()
	if err != nil {
		log.Println("bad user service")
		return nil, err
	}

	return &MasterClient{
		pc: pc,
		us: us,
	}, nil
}

// todo: move this functionality elsewhere. masterclient should only deal with the prestigechain
func (mc *MasterClient) AddNewUser(username, password string) error {
	return mc.us.AddNewUser(username, password)
}

// todo: move this functionality elsewhere. masterclient should only deal with the prestigechain
func (mc *MasterClient) UserExists(username string) bool {
	return mc.us.UserExists(username)
}

// todo: move this functionality elsewhere. masterclient should only deal with the prestigechain
func (mc *MasterClient) ValidateUserPassword(username, password string) error {
	return mc.us.ValidateUserPassword(username, password)
}

// Assumes user has already been validated
func (mc *MasterClient) AddNewAchievementTransaction(user string, value int, reason string, relevantTransactionIds [][]byte) error {
	tx := prestigechain.NewAchievementTX(user, value, reason, relevantTransactionIds)
	return mc.pc.AddBlock([]*prestigechain.Transaction{tx})
}

// Not implemented. MasterClient contains the master Prestigechain so no updates needed.
func (mc *MasterClient) FetchUpdates() error {
	return nil
}

func (mc *MasterClient) GetPrestigeChain() *prestigechain.Prestigechain {
	return mc.pc
}

//todo: consider caching
// Gets blocks from start (most recent) to end (least recent) exclusive of end.
// The newest block is at index 0 and the oldest block is at index n where length of the entire Prestigechain is n.
func (mc *MasterClient) GetBlocks(start, end int) []*prestigechain.PrestigeBlock {
	iterator := mc.pc.Iterator()
	numBlocks := end - start
	var blocks = make([]*prestigechain.PrestigeBlock, 50, numBlocks)
	for i := 0; i < end; i++ {
		if i < start {
			continue
		}

		block, _ := iterator.Next()
		if block == nil {
			blocks = blocks[0 : i-(start+1)]
			break
		}

		blocks[i-start] = block
	}

	return blocks
}

func DeleteMasterClient(mc *MasterClient) {
	prestigechain.DeletePrestigechain(mc.pc)
	DeleteUserService(mc.us)
}
