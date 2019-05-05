package client

import (
	"context"
	"log"
	"sync"

	"github.com/peterxu30/cloudchain"

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
		return nil, err
	}

	return &MasterClient{
		pc: pc,
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
// If more blocks are requested than exist in the PrestigeChain, GetBlocks will return all blocks in the PrestigeChain.
// The newest block is at index 0 and the oldest block is at index n where length of the entire Prestigechain is n.
func (mc *MasterClient) GetBlocks(ctx context.Context, numBlocks int) ([]*prestigechain.PrestigeBlock, error) {
	iterator, err := mc.pc.Iterator(ctx)
	if err != nil {
		return nil, err
	}

	var blocks = make([]*prestigechain.PrestigeBlock, 0, numBlocks)
	for i := 0; i < numBlocks; i++ {
		block, err := iterator.Next()
		if err == nil {
			blocks = append(blocks, block)
			continue
		} else if _, ok := err.(*cloudchain.StopIterationError); ok {
			break
		} else {
			return nil, err
		}
	}

	return blocks, nil
}

func DeleteMasterClient(ctx context.Context, mc *MasterClient) {
	prestigechain.DeletePrestigechain(ctx, mc.pc)
	DeleteUserService(mc.us)
}
