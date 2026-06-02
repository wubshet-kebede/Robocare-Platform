// package ws

// import (
// 	"log"
// 	"sync"
// )

// type Manager struct {
// 	clients map[*Client]bool
// 	mu      sync.RWMutex
// }

// func NewManager() *Manager {
// 	return &Manager{
// 		clients: make(map[*Client]bool),
// 	}
// }
// func (m *Manager) AddClient(c *Client) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	m.clients[c] = true

// 	log.Printf("Client connected → hospital=%s role=%s robot=%s",
// 		c.HospitalID, c.Role, c.RobotID)
// }

// func (m *Manager) RemoveClient(c *Client) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	delete(m.clients, c)
// 	close(c.Send)

// 	log.Printf("Client disconnected → hospital=%s", c.HospitalID)
// }
// func (m *Manager) BroadcastVitals(hospitalID string, data []byte) {
// 	m.mu.RLock()
// 	defer m.mu.RUnlock()

// 	for client := range m.clients {

// 		if client.HospitalID != hospitalID {
// 			continue
// 		}

// 		select {
// 		case client.Send <- data:
// 		default:
// 			log.Println("dropping slow client")
// 		}
// 	}
// }
// func (m *Manager) RouteMessage(sender *Client, msg Message, raw []byte) {

// 	switch msg.Type {

// 	// =========================
// 	// EXISTING SYSTEM
// 	// =========================
// 	case "vitals":
// 		m.BroadcastVitals(sender.HospitalID, raw)

// 	// =========================
// 	// TELEPRESENCE (NEW)
// 	// =========================
// 	case "offer", "answer", "candidate":

// 		m.mu.RLock()
// 		defer m.mu.RUnlock()

// 		for client := range m.clients {

// 			// match robot session
// 			if client.RobotID == msg.RobotID &&
// 				client != sender {

// 				select {
// 				case client.Send <- raw:
// 				default:
// 					log.Println("telepresence client slow")
// 				}
// 			}
// 		}
// 	case "register_robot":

//         sender.RobotID = msg.RobotID

//         log.Printf(
//         "Robot registered -> %s",
//         sender.RobotID,
//     )

//     case "register_session":
//          sender.HospitalID = msg.HospitalID
//          sender.Role = "doctor"
//          log.Println("Session registered:", sender.HospitalID)
// 	}

// }
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

func (m *Manager) RouteMessage(sender *Client, msg Message, raw []byte) {
	switch msg.Type {

	// =========================
	// EXISTING SYSTEM
	// =========================
	case "vitals":
		m.BroadcastVitals(sender.HospitalID, raw)

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

			// 🔥 DIRECT ROUTE: If the other client's RobotID matches the target ID, route it!
			if client.RobotID == msg.RobotID {
				select {
				case client.Send <- raw:
					routed = true
				default:
					log.Printf("Slow connection client dropped for robot: %s", msg.RobotID)
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
