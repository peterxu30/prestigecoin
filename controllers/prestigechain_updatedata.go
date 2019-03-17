package controllers

import "github.com/peterxu30/prestigecoin/prestigechain"

type PrestigechainUpdateData struct {
	User                   string
	Type                   prestigechain.TXType // not currently used
	Value                  int
	Reason                 string
	RelevantTransactionIds [][]byte
}
