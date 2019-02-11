package main

import (
	"fmt"
	"net/http"

	"github.com/peterxu30/prestigecoin/prestigechain"
)

type MasterClientController struct{}

func NewMasterClientController() *MasterClientController {
	return &MasterClientController{}
}

func (mcc *MasterClientController) handlePrestigechainInit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pcClient, err := prestigechain.NewMasterClient()
		if err != nil {
			fmt.Fprintf(w, "Prestigechain creation failed: %v", err)
			return
		}
		s.pcClient = pcClient

		fmt.Fprintf(w, "Prestigechain created.")
	}
}

func (mcc *MasterClientController) handlePrestigechainUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//s.pcClient.AddNewAchievementTransaction()
	}
}
