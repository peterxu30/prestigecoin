package main //change package to server-esque package

import (
	"fmt"
	"log"
	"net/http"

	prestigechain "github.com/peterxu30/prestigecoin/prestigechain"
)

type iServer interface {
	routes()
}

type server struct {
	pcClient *prestigechain.LocalClient
}

func NewServer() *server {
	return &server{}
}

func (s *server) ListenAndServe() {
	fmt.Printf("Prestigechain server started.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *server) handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Prestigechain - By the pChild")
	}
}

func (s *server) handlePrestigechainInit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pcClient := prestigechain.NewLocalClient()
		s.pcClient = pcClient

		fmt.Fprintf(w, "Prestigechain created.")
	}
}

func (s *server) handlePrestigechainUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body

		s.pcClient.AddNewAchievementTransaction()
	}
}

func getJsonParameters()

func main() {
	s := NewServer()
	s.routes()
	s.ListenAndServe()
}
