package mqttclient

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/admission"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/room"
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
    // check if there is an active admission for the patient in this room and get the patient id
    patient, err := admission.GetPatientByRoomID(body.RoomID)
    if err != nil {
	    http.Error(w, "patient not found in this room", http.StatusNotFound)
	    return
    } 
    
     goal := GoalPayload{
        HospitalID:   hospitalID.String(),
        PatientID:    patient.ID.String(),
        RobotID:      body.RobotID,
        TargetRoomID: body.RoomID.String(),
        DoctorID:     patient.AssignedDoctorID.String(),
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
