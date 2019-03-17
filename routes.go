package main //change package to server-esque package
import (
	"log"
	"net/http"

	"github.com/peterxu30/prestigecoin/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Initialize controllers here and add their methods to routes
var _masterClientController = controllers.NewMasterClientController()

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handleDefault(),
	},
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
