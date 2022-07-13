package http

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"

	chat "ws-chat.example"
)

func (s Service) newClient(conn *websocket.Conn) error {
	c := &client{
		rwc: conn,
		p:   s.p,
	}

	c.serve()
	return nil
}

type client struct {
	// id  string

	rwc *websocket.Conn
	p   *pool
}

func (c *client) close() error {
	c.p.unr <- c
	return c.rwc.Close()
}

func (cli client) read() (messageType int, p []byte, err error) { return cli.rwc.ReadMessage() }

func (cli client) write(v any) error { return cli.rwc.WriteJSON(v) }

func (cli *client) serve() {
	defer cli.close()

	for {
		messageType, p, err := cli.read()
		if err != nil {
			log.Println(err)
			return
		}
		msg := chat.Message{Type: messageType, Body: string(p)}
		cli.p.bc <- msg
		fmt.Printf("Message Received: %+v\n", msg)
	}
}

// func (c client) serve() {}

type pool struct {
	mu   sync.RWMutex
	r    chan *client
	unr  chan *client
	clis map[*client]bool
	bc   chan chat.Message
}

func newPool() *pool {
	p := &pool{
		r:    make(chan *client),
		unr:  make(chan *client),
		bc:   make(chan chat.Message),
		clis: make(map[*client]bool),
	}
	go p.run()
	return p
}

func (p *pool) broadcast(v any) {
	for cli := range p.clis {
		cli.write(v)
	}
}

func (p *pool) add(cli *client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.clis[cli] = true
	fmt.Println("Size of Connection Pool: ", len(p.clis))
}

func (p *pool) remove(cli *client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.clis, cli)
}

func (p *pool) run() {
	for {
		select {
		case cli := <-p.r:
			p.add(cli)
			p.broadcast(chat.Message{Type: 1, Body: "New User Joined..."})
		case cli := <-p.unr:
			p.remove(cli)
			p.broadcast((chat.Message{Type: 1, Body: "User Disconnected..."}))
		case msg := <-p.bc:
			p.broadcast(msg)
		}
	}
}
