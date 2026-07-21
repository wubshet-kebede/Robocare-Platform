package ws

import (
	"encoding/json"
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

	log.Printf("Client connected → hospital=%s role=%s robot=%s",
		c.HospitalID, c.Role, c.RobotID)
}

func (m *Manager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, c)
	close(c.Send)

	log.Printf("Client disconnected → hospital=%s", c.HospitalID)
}

func (m *Manager) BroadcastVitals(hospitalID string,robotID string, data []byte) {
	 log.Printf("BroadcastVitals called for hospital=%s robot=%s", hospitalID, robotID)
	 log.Printf("Connected clients: %d", len(m.clients))
	m.mu.RLock()
	defer m.mu.RUnlock()
	type WSVitalsEnvelope struct {
		Type    string      `json:"type"`
		RobotID string      `json:"robotId"`
		Data    interface{} `json:"data"`
	}
	var rawPayload interface{}
	json.Unmarshal(data, &rawPayload)

	envelope := WSVitalsEnvelope{
		Type:    "vitals",
		RobotID: robotID,
		Data:    rawPayload,
	}
	envelopeBytes, _ := json.Marshal(envelope)

	for client := range m.clients {

		isMatchingHospital := (client.HospitalID != "" && client.HospitalID == hospitalID)
		isMatchingRobot := (client.RobotID != "" && client.RobotID == robotID)

		if isMatchingHospital || isMatchingRobot {
			select {
			case client.Send <- envelopeBytes:
				
			default:
				
			}
		} 
	}
}

func (m *Manager) RouteMessage(sender *Client, msg Message, raw []byte) {
	switch msg.Type {

	// =========================
	// EXISTING SYSTEM
	// =========================
	case "vitals":
		m.BroadcastVitals(sender.HospitalID, sender.RobotID, raw)

	// =========================
	// TELEPRESENCE (FIXED ROUTING)
	// =========================
	case "offer", "answer", "candidate":
		m.mu.RLock()
		defer m.mu.RUnlock()

		routed := false

		for client := range m.clients {
			// Skip sending the message back to ourselves
			if client == sender {
				continue
			}

			// If the other client's RobotID matches the target ID, route it!
			if client.RobotID == msg.RobotID {
				select {
				case client.Send <- raw:
					routed = true
				default:
					
				}
			}
		}

		if !routed {
			log.Printf("[Warning] Telepresence message %s for robot %s could not find matching peer", msg.Type, msg.RobotID)
		}
	case "register_robot":
		m.mu.Lock()
		sender.RobotID = msg.RobotID
		m.mu.Unlock()
		log.Printf("Robot registered -> %s", sender.RobotID)

	case "register_session":
		m.mu.Lock()
		sender.HospitalID = msg.HospitalID
		sender.Role = "doctor"
		m.mu.Unlock()
		log.Println("Session registered:", sender.HospitalID)
	}
}
