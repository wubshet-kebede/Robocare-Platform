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
	log.Printf("Client connected → hospital=%s patient=%s",
		c.HospitalID, c.PatientID)
}

func (m *Manager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, c)
	log.Printf("Client disconnected → hospital=%s patient=%s",
		c.HospitalID, c.PatientID)
}
func (m *Manager) BroadcastVitals(hospitalID string, patientID string, data []byte) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for client := range m.clients {

		// 🔐 FILTER BY TENANT (hospital)
		if client.HospitalID != hospitalID {
			continue
		}

		// 🎯 FILTER BY PATIENT (if subscribed)
		if client.PatientID != "" && client.PatientID != patientID {
			continue
		}

		err := client.Conn.WriteMessage(1, data)
		if err != nil {
			log.Println("WS send error:", err)
			client.Conn.Close()
			delete(m.clients, client)
		}
	}
}