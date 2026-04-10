package mqttclient

import (
	"sync"
	"time"

	"github.com/google/uuid"
)
type GoalPayload struct {
    GoalID       string    `json:"goal_id"`
    MapID        string    `json:"map_id"`
	RobotID     string    `json:"robot_id"`
    TargetRoomID string    `json:"target_room_id"`
    X  float64   `json:"x"`
	Y  float64   `json:"y"`
	Theta  float64   `json:"theta"`
    DoctorID     string    `json:"doctor_id"`
    Timestamp    time.Time `json:"timestamp"`
    HMAC         string    `json:"hmac,omitempty"`
}

type MQTTPublisher struct {
    clients map[string]*MQTTClientManager // hospitalUUID -> client
    secrets map[string]string             // hospitalUUID -> HMAC secret
    broker  string                         // MQTT broker URL
    port    int                            // MQTT port
	mu      sync.Mutex                     // protects clients map

}
type PublishGoalRequest struct {
    RobotID string    `json:"robot_id"`
    RoomID  uuid.UUID `json:"room_id"`
}