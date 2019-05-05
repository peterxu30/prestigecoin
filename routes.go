package main //change package to server-esque package

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/peterxu30/prestigecoin/controllers"
)

const (
	prestigecoinProjectId = "prestigecoin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Initialize controllers here and add their methods to routes
var _masterClientController = controllers.NewMasterClientController(context.Background(), prestigecoinProjectId)

// Consider moving this to some controller as well
func handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Prestigechain - By the pChild")
	}
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handleDefault(),
	},
	// Route{
	// 	"Init",
	// 	"GET",
	// 	"/init",
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		// defer func() {
	// 		// 	if re := recover(); re != nil {
	// 		// 		fmt.Fprintln(w, "Init failed", re)
	// 		// 	}
	// 		// }()
	// 		msg := _masterClientController.Init(r.Context(), prestigecoinProjectId)
	// 		fmt.Fprintln(w, msg)
	// 	},
	// },
	// Route{
	// 	"AddNewUser",
	// 	"POST",
	// 	"/newuser",
	// 	s.handleAddNewUser(),
	// },
	Route{
		"GetBlocks",
		"GET",
		"/getblocks",
		_masterClientController.HandlePrestigechainGet(),
	},
	Route{
		"AddBlock",
		"POST",
		"/addblock",
		_masterClientController.HandlePrestigechainUpdate(),
	},
}

// Add all routes from all controllers here
func InitializeRoutes(s *server) {
	for _, route := range routes {
		var handler http.HandlerFunc
		log.Println(route.Name)
		handler = route.HandlerFunc

		s.router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}
