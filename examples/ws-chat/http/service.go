package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"

	"ws-chat.example/http/ws"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/http context value " + k.name }

type Service struct {
	m *chi.Mux
	// ps map[*pool]bool
	p *ws.Pool
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.m.ServeHTTP(w, r) }

func NewService() *Service {
	s := &Service{
		m: chi.NewMux(),
		// ps: make(map[*pool]bool),
		p: ws.NewPool(),
	}
	s.routes()
	return s
}

func (s *Service) newClient(w http.ResponseWriter, r *http.Request) error {
	return ws.NewClient(s.p, r.Context().Value(upgradeKey).(*websocket.Conn)).Serve()
}
