package http

import (
	"net/http"

	"github.com/gorilla/websocket"
	t "github.com/topheruk-go/util/template"
	h "ws-chat.example/pkg/http"
)

func (s *Service) routes() {
	s.m.Handle("/assets/*", h.FileServer("/assets", "assets"))
	s.m.Get("/", s.handleIndex())
	s.m.HandleFunc("/chat", chain(s.handleChat(), s.upgradeHTTP()))
}

func (s Service) handleIndex() http.HandlerFunc {
	render, err := t.Render("pages/index.html")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, nil)
	}
}

func (s Service) handleChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn := r.Context().Value(upgradeKey).(*websocket.Conn)

		// create new client and serve
		go s.newClient(conn)
	}
}
