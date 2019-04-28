package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/peterxu30/prestigecoin/client"
	"github.com/peterxu30/prestigecoin/prestigechain"
)

type MasterClientController struct {
	client *client.MasterClient
}

func NewMasterClientController(ctx context.Context, projectId string) *MasterClientController {
	mcc := &MasterClientController{}
	mcc.client = client.GetOrCreateMasterClient(ctx, projectId)
	return mcc
}

func (mcc *MasterClientController) HandlePrestigechainUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var updateData PrestigechainUpdateData
		err := decoder.Decode(&updateData)
		if err != nil {
			// to do: better failure logic
			log.Println(err)
		}

		if updateData.Type == prestigechain.Achievement {
			mcc.client.AddNewAchievementTransaction(r.Context(), updateData.User, updateData.Value, updateData.Reason, updateData.RelevantTransactionIds)
			log.Printf("Block added for user %s with value %v for reason %s", updateData.User, updateData.Value, updateData.Reason)
		}
	}
}

func (mcc *MasterClientController) HandlePrestigechainGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start, end := 0, 50

		startParam, ok := r.URL.Query()["start"]
		if ok && len(startParam[0]) >= 1 {
			if startIndex, err := strconv.Atoi(startParam[0]); err == nil {
				start = startIndex
			}
		}

		endParam, ok := r.URL.Query()["end"]
		if ok && len(endParam[0]) >= 1 {
			if endIndex, err := strconv.Atoi(endParam[0]); err == nil {
				end = endIndex
			}
		}

		blocks, err := mcc.client.GetBlocks(r.Context(), start, end)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		transactions := make([]*prestigechain.Transaction, 0, len(blocks))
		for _, block := range blocks {
			transactions = append(transactions, block.Transactions[0])
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(blocks)
	}
}
