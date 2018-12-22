package main //change package to server-esque package
import "net/http"

func (s *server) routes() {
	http.HandleFunc("/", s.handleDefault())
	http.HandleFunc("/init", s.handlePrestigechainInit())
}
