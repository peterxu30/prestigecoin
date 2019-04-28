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

func main() {
	s := NewServer()
	s.ListenAndServe()
}
