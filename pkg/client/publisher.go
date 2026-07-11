package client

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Publisher[T any] struct {
	Connector *websocket.Conn
	Address   string
}

func (p *Publisher[T]) Publish(msg T) error {
	marshal, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = p.Connector.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		return err
	}

	return nil
}
