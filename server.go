package main //change package to server-esque package

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/peterxu30/prestigecoin/client"
)

type iServer interface {
	routes()
}

type server struct {
	pcClient *client.MasterClient
	router   *mux.Router
}

func NewServer() *server {
	server := &server{
		router: mux.NewRouter().StrictSlash(true),
	}

	InitializeRoutes(server)

	return server
}

func (s *server) ListenAndServe() {
	log.Print("Prestigechain server started.")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(s.router)))
}

// Consider moving this to some controller as well
func handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Prestigechain - By the pChild")
	}
}

// ---- Move below functionality to User Service Controller ----

// func (s *server) handleAddNewUser() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//all temp code
// 		var userData prestigechain.UserData
// 		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		log.Println(body)

// 		if err := json.Unmarshal(body, &userData); err != nil { // unmarshall body contents as a type Candidate
// 			w.WriteHeader(422) // unprocessable entity
// 			log.Println(err)
// 			if err := json.NewEncoder(w).Encode(err); err != nil {
// 				log.Fatalln("Error AddProduct unmarshalling data", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 				return
// 			}
// 		}

// 		log.Println(userData.Username + " " + userData.Password)

// 		fmt.Fprintf(w, "User added")
// 	}
// }

func main() {
	s := NewServer()
	s.ListenAndServe()
}
