package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all (OK for now, restrict in production)
	},
}

type Handler struct {
	manager *Manager
}

func NewHandler(m *Manager) *Handler {
	return &Handler{manager: m}
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	hospitalID := r.URL.Query().Get("hospital_id")
	patientID := r.URL.Query().Get("patient_id")

	if hospitalID == "" {
		http.Error(w, "hospital_id is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	client := &Client{
		Conn:       conn,
		HospitalID: hospitalID,
		PatientID:  patientID,
		Send:       make(chan []byte, 256),
	}
	h.manager.AddClient(client)
    go h.writePump(client)
	go h.listen(client)
}

func (h *Handler) listen(c *Client) {
	defer func() {
		h.manager.RemoveClient(c)
		c.Conn.Close()
	}()

	for {
	
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
func (h *Handler) writePump(c *Client) {
	defer func() {
		c.Conn.Close()
	}()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}