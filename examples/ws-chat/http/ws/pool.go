package ws

import (
	"fmt"
	"log"
	"sync"

	chat "ws-chat.example"
)

type Pool struct {
	mu sync.RWMutex

	r   chan *Client
	unr chan *Client
	bc  chan chat.Message

	clis map[*Client]bool
}

func NewPool() *Pool {
	p := &Pool{
		r:    make(chan *Client),
		unr:  make(chan *Client),
		bc:   make(chan chat.Message),
		clis: make(map[*Client]bool),
	}
	go p.run()
	return p
}

func (p *Pool) broadcast(v any) {
	for cli := range p.clis {
		cli.Write(v)
	}
}

func (p *Pool) add(cli *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.clis[cli] = true
	fmt.Println("Size of Connection Pool: ", len(p.clis))
	go cli.Serve()
}

func (p *Pool) remove(cli *Client) {
	p.mu.Lock()
	defer func() {
		p.mu.Unlock()
		log.Println("closing")
		cli.Close()
	}()

	delete(p.clis, cli)
}

func (p *Pool) run() {
	for {
		select {
		case cli := <-p.r:
			p.add(cli)
			p.broadcast("New User Joined...")
		case cli := <-p.unr:
			p.remove(cli)
			p.broadcast("User Disconnected...")
		case msg := <-p.bc:
			p.broadcast(msg)
		}
	}
}
