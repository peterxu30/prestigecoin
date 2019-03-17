package client

import (
	"github.com/peterxu30/prestigecoin/prestigechain"
)

type Client interface {
	AddNewAchievementTransaction(username, value int, reason string, relevantTxIds [][]byte) error
	FetchUpdates() error
	GetPrestigechain() *prestigechain.Prestigechain
}
