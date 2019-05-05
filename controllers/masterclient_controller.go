package controllers

import (
	"context"
	"encoding/json"
	"fmt"
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

//TEST METHOD
// func (mcc *MasterClientController) Init(ctx context.Context, projectId string) (msg string) {
// 	defer func() {
// 		if re := recover(); re != nil {
// 			msg = string(debug.Stack())
// 		}
// 	}()
// 	mcc.client = client.GetOrCreateMasterClient(ctx, projectId)
// 	return ""
// }

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
		numBlocks := 50

		numBlocksParam, ok := r.URL.Query()["numBlocks"]
		if ok && len(numBlocksParam[0]) >= 1 {
			if numBlocksValue, err := strconv.Atoi(numBlocksParam[0]); err == nil {
				numBlocks = numBlocksValue
			} else {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
				return
			}
		}

		// errMsg = "No error"
		// defer func() {
		// 	if re := recover(); re != nil {
		// 		errMsg = string(debug.Stack())
		// 	}
		// }()

		blocks, err := mcc.client.GetBlocks(r.Context(), numBlocks)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		transactions := make([]*prestigechain.Transaction, 0, len(blocks))
		for _, block := range blocks {
			transactions = append(transactions, block.Transactions[0])
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transactions)
	}
}
