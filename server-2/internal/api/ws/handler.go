package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ⚠️ allow all (OK for now, restrict in production)
	},
}

type Handler struct {
	manager *Manager
}

func NewHandler(m *Manager) *Handler {
	return &Handler{manager: m}
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	// 🔥 1. Extract query params
	hospitalID := r.URL.Query().Get("hospital_id")
	patientID := r.URL.Query().Get("patient_id")

	if hospitalID == "" {
		http.Error(w, "hospital_id is required", http.StatusBadRequest)
		return
	}

	// 🔥 2. Upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	// 🔥 3. Create client
	client := &Client{
		Conn:       conn,
		HospitalID: hospitalID,
		PatientID:  patientID,
	}

	// 🔥 4. Register client
	h.manager.AddClient(client)

	// 🔥 5. Listen for disconnect
	go h.listen(client)
}

func (h *Handler) listen(c *Client) {
	defer func() {
		h.manager.RemoveClient(c)
		c.Conn.Close()
	}()

	for {
		// We don’t care about messages yet
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
	}
}