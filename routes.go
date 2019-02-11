package main //change package to server-esque package
import (
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		s.handleDefault(),
	},
	Route{
		"Init",
		"GET",
		"/init",
		s.handlePrestigechainInit(),
	},
	Route{
		"AddNewUser",
		"POST",
		"/newuser",
		s.handleAddNewUser(),
	},
}

// Add all routes from all controllers here
func InitializeRoutes(s *server) Routes {
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
