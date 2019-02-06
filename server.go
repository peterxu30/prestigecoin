package main //change package to server-esque package

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	prestigechain "github.com/peterxu30/prestigecoin/prestigechain"
)

type iServer interface {
	routes()
}

type server struct {
	pcClient *prestigechain.MasterClient
	router   *mux.Router
}

func NewServer() *server {
	server := &server{
		router: mux.NewRouter().StrictSlash(true),
	}

	InitializeRouter(server)

	return server
}

func (s *server) ListenAndServe() {
	log.Print("Prestigechain server started.")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(s.router)))
}

func (s *server) handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Prestigechain - By the pChild")
	}
}

// ---- Move below functionality to Prestigechain Controller ----

func (s *server) handlePrestigechainInit() http.HandlerFunc {
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

func (s *server) handlePrestigechainUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//s.pcClient.AddNewAchievementTransaction()
	}
}

// ---- Move below functionality to User Service Controller ----

func (s *server) handleAddNewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//all temp code
		var userData prestigechain.UserData
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(body)

		if err := json.Unmarshal(body, &userData); err != nil { // unmarshall body contents as a type Candidate
			w.WriteHeader(422) // unprocessable entity
			log.Println(err)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Fatalln("Error AddProduct unmarshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		log.Println(userData.Username + " " + userData.Password)

		fmt.Fprintf(w, "User added")
	}
}

func main() {
	s := NewServer()
	s.ListenAndServe()
}
