package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/poll"
)

type Server struct {
	pollController *poll.Controller
	router         *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(storer poll.Storer) *Server {
	return &Server{
		pollController: poll.NewController(storer),
		router:         chi.NewRouter(),
	}
}
