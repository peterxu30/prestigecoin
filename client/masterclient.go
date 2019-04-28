package client

import (
	"context"
	"log"
	"sync"

	"github.com/peterxu30/prestigecoin/prestigechain"
)

// TODO: Make the MasterClient thread safe.

// MasterClient contains the master Prestigechain.
type MasterClient struct {
	pc *prestigechain.Prestigechain
	us *UserService
}

var _masterClient *MasterClient
var once sync.Once

func GetOrCreateMasterClient(ctx context.Context, projectId string) *MasterClient {
	once.Do(func() {
		var err error
		_masterClient, err = newMasterClient(ctx, projectId)
		if err != nil {
			log.Panicln("Failed to create MasterClient", err)
		}
	})

	return _masterClient
}

func newMasterClient(ctx context.Context, projectId string) (*MasterClient, error) {
	pc, err := prestigechain.NewPrestigechain(ctx, projectId)
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

// Assumes user has already been validated
func (mc *MasterClient) AddNewAchievementTransaction(ctx context.Context, user string, value int, reason string, relevantTransactionIds [][]byte) (*prestigechain.PrestigeBlock, error) {
	tx := prestigechain.NewAchievementTX(user, value, reason, relevantTransactionIds)
	return mc.pc.AddBlock(ctx, []*prestigechain.Transaction{tx})
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
func (mc *MasterClient) GetBlocks(ctx context.Context, start, end int) ([]*prestigechain.PrestigeBlock, error) {
	iterator, err := mc.pc.Iterator(ctx)
	if err != nil {
		return nil, err
	}

	numBlocks := end - start
	var blocks = make([]*prestigechain.PrestigeBlock, 0, numBlocks)
	for i := 0; i < end; i++ {
		block, _ := iterator.Next()
		if block == nil {
			break
		}

		if i >= start {
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func DeleteMasterClient(ctx context.Context, mc *MasterClient) {
	prestigechain.DeletePrestigechain(ctx, mc.pc)
	DeleteUserService(mc.us)
}
