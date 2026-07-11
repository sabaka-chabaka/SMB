package broker

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	broker *Broker
	conn   *websocket.Conn
	send   chan []byte
}

func NewClient(b *Broker, conn *websocket.Conn) *Client {
	return &Client{
		broker: b,
		conn:   conn,
		send:   make(chan []byte),
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.broker.Unregister(c)
		err := c.conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.broker.Broadcast(message)
	}
}

func (c *Client) WritePump() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}
