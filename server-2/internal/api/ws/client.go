package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn       *websocket.Conn
	HospitalID string
	Send       chan []byte
}