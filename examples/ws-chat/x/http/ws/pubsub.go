package ws

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Typ  string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type MessageType int

func (msg MessageType) String() string { return "unknown" }

const (
	Unknown MessageType = iota
)

type Handler interface {
	ServeWS(sub *Subcriber)
}

type HandlerFunc func(sub *Subcriber)

func (f HandlerFunc) ServeWS(sub *Subcriber) { f(sub) }

type Publisher struct {
	mu sync.RWMutex

	msub map[string]subcription

	ss map[*Subcriber]bool
}

func (pub *Publisher) handler(sub *Subcriber) (h Handler, m string) {
	pub.mu.Lock()
	defer pub.mu.Unlock()

	if h == nil {
		h, m = pub.match(sub.Msg.Typ)
	}

	if h == nil {
		h, m = nil, ""
	}

	return
}

func (pub *Publisher) Handle(msgTyp string, handler Handler) {
	pub.mu.Lock()
	defer pub.mu.Unlock()

	if msgTyp == "" {
		panic("ws: unknown message")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if _, exist := pub.msub[msgTyp]; exist {
		panic("http: multiple registrations for " + msgTyp)
	}

	if pub.msub == nil {
		pub.msub = make(map[string]subcription)
	}
	e := subcription{h: handler, msgTyp: msgTyp}
	pub.msub[msgTyp] = e
}

// sanitise the msg to lowercase
func (pub *Publisher) HandleFunc(msg string, handler func(sub *Subcriber)) {
	if handler == nil {
		panic("ws: nil handler")
	}

	pub.Handle(msg, HandlerFunc(handler))
}

func (pub *Publisher) ServeWS(sub *Subcriber) {
	h, _ := pub.handler(sub)
	for s := range pub.ss {
		h.ServeWS(s)
	}
}

type subcription struct {
	h      Handler
	msgTyp string
}

func (pub *Publisher) match(msgTyp string) (Handler, string) {
	v, ok := pub.msub[msgTyp]
	if ok {
		return v.h, v.msgTyp
	}

	return nil, ""
}

type Subcriber struct {
	Msg  Message
	Conn *websocket.Conn
	Pub  *Publisher
}

func newSubscriber() *Subcriber {
	var s *Subcriber
	go s.serve()
	return s
}

func (sub Subcriber) Close() error { return sub.Conn.Close() }

func (sub Subcriber) Read() error { return sub.Conn.ReadJSON(sub.Msg) }

func (sub Subcriber) Write() error { return sub.Conn.WriteJSON(&sub.Msg) }

func (sub *Subcriber) serve() {
	defer sub.Close()
	for {

		if err := sub.Read(); err != nil {
			// handle error
			return
		}

		sub.Pub.ServeWS(sub)
	}
}
