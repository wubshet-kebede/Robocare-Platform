package handler

// import (
// 	"net/http"

// 	wsHub "github.com/wubshet-kebede/robocare-platform/server-2/internal/api/websocket"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func TelepresenceWSHandler(w http.ResponseWriter, r *http.Request) {

// 	role := r.URL.Query().Get("role")

// 	if role == "" {
// 		http.Error(w, "role required", http.StatusBadRequest)
// 		return
// 	}

// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return
// 	}

// 	client := &wsHub.Client{
// 		Conn: conn,
// 		Role: role,
// 	}

// 	wsHub.TelepresenceHub.mu.Lock()

// 	if role == "doctor" {
// 		wsHub.TelepresenceHub.Doctor = client
// 	}

// 	if role == "robot" {
// 		wsHub.TelepresenceHub.Robot = client
// 	}

// 	wsHub.TelepresenceHub.mu.Unlock()

// 	for {

// 		_, message, err := conn.ReadMessage()

// 		if err != nil {
// 			break
// 		}

// 		wsHub.TelepresenceHub.mu.Lock()

// 		if role == "doctor" && wsHub.TelepresenceHub.Robot != nil {

// 			wsHub.TelepresenceHub.Robot.Conn.WriteMessage(
// 				websocket.TextMessage,
// 				message,
// 			)
// 		}

// 		if role == "robot" && wsHub.TelepresenceHub.Doctor != nil {

// 			wsHub.TelepresenceHub.Doctor.Conn.WriteMessage(
// 				websocket.TextMessage,
// 				message,
// 			)
// 		}

// 		wsHub.TelepresenceHub.mu.Unlock()
// 	}

// 	conn.Close()
// }