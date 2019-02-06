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

// Add all routes from all controllers here
func (s *server) initializeRoutes() Routes {
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

	return routes
}

func InitializeRouter(s *server) {
	routes := s.initializeRoutes()

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
