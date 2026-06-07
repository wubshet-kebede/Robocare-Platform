package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
type Handler struct {
	manager *Manager
}

func NewHandler(m *Manager) *Handler {
	return &Handler{manager: m}
}
// func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
// 	hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
// 	role, Rok := r.Context().Value(middleware.RoleKey).(string)

// 	if !ok || !Rok {
// 		http.Error(w, "missing hospital id or role", http.StatusBadRequest)
// 		return
// 	}
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade failed:", err)
// 		return
// 	}

// 	client := &Client{
// 		Conn:       conn,
// 		HospitalID: hospitalID.String(),
// 		Role:       role,
// 		RobotID:    " ",
// 		Send:       make(chan []byte, 256),
// 	}

// 	h.manager.AddClient(client)

// 	go h.writePump(client)
// 	go h.listen(client)
// }
func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 256),

		// unknown at connection time
		HospitalID: "",
		Role:       "",
		RobotID:    "",
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

		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			break
		}

		var m Message

		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Println("invalid message:", err)
			continue
		}

		h.manager.RouteMessage(c, m, msg)
	}
}
func (h *Handler) writePump(c *Client) {
    
	defer func() {
		c.Conn.Close()
	}()

	for msg := range c.Send {
		log.Printf(
        "writePump sending %d bytes",
        len(msg),
    )

		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
	
}