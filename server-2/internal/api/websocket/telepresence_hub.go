package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

//
// =========================
// MESSAGE FORMAT
// =========================
//

type SignalMessage struct {
	Type      string `json:"type"`      // offer | answer | candidate | register
	RobotID   string `json:"robotId"`
	SDP       string `json:"sdp,omitempty"`
	Candidate string `json:"candidate,omitempty"`
}

//
// =========================
// CLIENT STRUCT
// =========================
//

type Client struct {
	conn    *websocket.Conn
	role    string // "robot" | "doctor"
	robotID string
}

//
// =========================
// HUB (CORE ROUTER)
// =========================
//

type TelepresenceHub struct {
	robots  map[string]*Client
	doctors map[string]*Client
	mu      sync.RWMutex
}

func NewTelepresenceHub() *TelepresenceHub {
	return &TelepresenceHub{
		robots:  make(map[string]*Client),
		doctors: make(map[string]*Client),
	}
}

//
// =========================
// REGISTER CLIENT
// =========================
//

func (h *TelepresenceHub) register(client *Client) {

	h.mu.Lock()
	defer h.mu.Unlock()

	if client.role == "robot" {
		h.robots[client.robotID] = client
		log.Println("[Telepresence] Robot connected:", client.robotID)
		return
	}

	if client.role == "doctor" {
		h.doctors[client.robotID] = client
		log.Println("[Telepresence] Doctor connected:", client.robotID)
		return
	}
}

//
// =========================
// ROUTE MESSAGE
// =========================
//

func (h *TelepresenceHub) route(sender *Client, msg SignalMessage) {

	h.mu.RLock()
	defer h.mu.RUnlock()

	switch sender.role {

	case "doctor":

		robot := h.robots[msg.RobotID]
		if robot != nil {

			err := robot.conn.WriteJSON(msg)
			if err != nil {
				log.Println("Error sending to robot:", err)
			}
		}

	case "robot":

		doctor := h.doctors[msg.RobotID]
		if doctor != nil {

			err := doctor.conn.WriteJSON(msg)
			if err != nil {
				log.Println("Error sending to doctor:", err)
			}
		}
	}
}

//
// =========================
// WEBSOCKET UPGRADE
// =========================
//

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // IMPORTANT: lock this in production later
	},
}

//
// =========================
// MAIN HANDLER
// =========================
//

func (h *TelepresenceHub) HandleWS(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	role := r.URL.Query().Get("role")
	robotID := r.URL.Query().Get("robotId")

	client := &Client{
		conn:    conn,
		role:    role,
		robotID: robotID,
	}

	h.register(client)

	defer func() {
		h.mu.Lock()
		defer h.mu.Unlock()

		if client.role == "robot" {
			delete(h.robots, client.robotID)
		}
		if client.role == "doctor" {
			delete(h.doctors, client.robotID)
		}

		conn.Close()
		log.Println("[Telepresence] disconnected:", client.role, client.robotID)
	}()

	for {

		var msg SignalMessage

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("read error:", err)
			return
		}

		h.route(client, msg)
	}
}