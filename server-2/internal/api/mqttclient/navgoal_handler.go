package mqttclient

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
)

var publisher *MQTTPublisher

func SetMQTTPublisher(p *MQTTPublisher) {
    publisher = p
}

func PublishNavGoalHandler(w http.ResponseWriter, r *http.Request) {
    hospitalID, ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
    if !ok {
        http.Error(w, "missing hospital id in context", http.StatusBadRequest)
        return
    }

    var body GoalPayload
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    goal := GoalPayload{
        GoalID:       body.GoalID,
        MapID:        body.MapID,
        RobotID:      body.RobotID,
        TargetRoomID: body.TargetRoomID,
        X:            body.X,
        Y:            body.Y,
        Theta:        body.Theta,
        DoctorID:     body.DoctorID,
        Timestamp:    time.Now(),
    }

    err := publisher.PublishNavGoal(hospitalID.String(), body.RobotID, goal)
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
