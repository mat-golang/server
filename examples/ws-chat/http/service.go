package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/http context value " + k.name }

type Service struct {
	m *chi.Mux
	// ps map[*pool]bool
	p *pool
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.m.ServeHTTP(w, r) }

func NewService() *Service {
	s := &Service{
		m: chi.NewMux(),
		// ps: make(map[*pool]bool),
		p: newPool(),
	}
	s.routes()
	return s
}
