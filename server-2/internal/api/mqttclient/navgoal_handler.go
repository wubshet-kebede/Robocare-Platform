package mqttclient

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/room"
)

var publisher *MQTTPublisher

func SetMQTTPublisher(p *MQTTPublisher) {
    publisher = p
}

func PublishNavGoalHandler(w http.ResponseWriter, r *http.Request) {
    hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
    userID, uok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)

    if !ok || !uok {
        http.Error(w, "missing hospital id or user id in context", http.StatusBadRequest)
        return
    }
    
    var body PublishGoalRequest
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    } 
    
    room, err := room.GetRoomByID(body.RoomID )
    if err != nil {
        http.Error(w, "room not found", http.StatusNotFound)
        return
    }

     goal := GoalPayload{
        HospitalID:   hospitalID.String(),
        PatientID:    body.PatientID,
        RobotID:      body.RobotID,
        TargetRoomID: body.RoomID.String(),
        DoctorID:     userID.String(),
        X:            room.X,
        Y:            room.Y,
        Theta:        room.Yaw,
        Timestamp:    time.Now(),
    }

    err = publisher.PublishNavGoal(hospitalID.String(), body.RobotID, goal)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "status":  "error",
            "message": err.Error(),
        })
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "success",
        "message": "Goal published to robot successfully",
    })
}
