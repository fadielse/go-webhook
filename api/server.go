package api

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/xcloud-webhook", s)
}

func (s *Server) xcloudToDiscord() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(r.Body))
	}
}