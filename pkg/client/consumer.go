package client

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Consumer[T any] struct {
	Connector *websocket.Conn
	Address   string
	Callback  func(T)
}

func (c *Consumer[T]) Consume() {
	for {
		_, payload, err := c.Connector.ReadMessage()
		if err != nil {
			panic(err)
		}

		c.processData(payload)
	}
}

func (c *Consumer[T]) processData(data []byte) {
	var msg T

	err := json.Unmarshal(data, &msg)
	if err != nil {
		return
	}

	c.Callback(msg)
}
