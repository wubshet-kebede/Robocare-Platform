package ws

import (
	"log"
	"sync"
)

type Manager struct {
	clients map[*Client]bool
	mu      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[*Client]bool),
	}
}

func (m *Manager) AddClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[c] = true
	log.Printf("Client connected → hospital=%s",
		c.HospitalID)
}

func (m *Manager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, c)
	close(c.Send) 
	log.Printf("Client disconnected → hospital=%s ",
		c.HospitalID, )
}
func (m *Manager) BroadcastVitals(hospitalID string, data []byte) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for client := range m.clients {

	
		if client.HospitalID != hospitalID {
			continue
		}
		select {
       case client.Send <- data:

       default:
	
	log.Println("dropping slow client")
}
		
	}
}