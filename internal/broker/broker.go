package broker

import "sync"

type Broker struct {
	clients map[*Client]bool
	mu      sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{clients: make(map[*Client]bool)}
}

func (b *Broker) Register(c *Client) {
	b.mu.Lock()
	b.clients[c] = true
	b.mu.Unlock()
}

func (b *Broker) Unregister(c *Client) {
	b.mu.Lock()
	if _, ok := b.clients[c]; ok {
		delete(b.clients, c)
		close(c.send)
	}
	b.mu.Unlock()
}

func (b *Broker) Broadcast(message []byte) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for client := range b.clients {
		select {
		case client.send <- message:
		default:

		}
	}
}
