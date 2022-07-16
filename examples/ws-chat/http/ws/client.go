package ws

import (
	"fmt"

	"github.com/gorilla/websocket"

	chat "ws-chat.example"
)

func NewClient(p *Pool, rwc *websocket.Conn) *Client {
	c := &Client{rwc, p}
	p.r <- c
	return c
}

type Client struct {
	rwc *websocket.Conn
	p   *Pool
}

func (c *Client) Close() error {
	c.p.unr <- c
	return c.rwc.Close()
}

func (cli Client) Read(v any) error { return cli.rwc.ReadJSON(v) }

func (cli Client) Write(v any) error { return cli.rwc.WriteJSON(v) }

func (cli *Client) Serve() error {

	for {
		var msg chat.Message
		if err := cli.Read(&msg); err != nil {
			cli.p.unr <- cli
			return err
		}

		// doSomething with message

		cli.p.bc <- msg
		fmt.Printf("Message Received: %+v\n", msg)
	}
}

// TODO -- if error, close connection
