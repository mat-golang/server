package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Service struct {
	m *chi.Mux
}

func NewService() *Service {
	s := &Service{chi.NewMux()}
	s.routes()
	return s
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.m.ServeHTTP(w, r) }
