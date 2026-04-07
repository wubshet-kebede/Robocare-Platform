package robot

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/middleware"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/robot"
)

type RobotReq struct {
    SerialNumber string `json:"serial_number"`
    Model        string `json:"model"`
}

func CreateRobotHandler(w http.ResponseWriter, r *http.Request) {
     hospitalID , ok := r.Context().Value(middleware.HospitalIDKey).(uuid.UUID)
	 if !ok  {
		http.Error(w, "missing hosppital id", http.StatusBadRequest)
		return
	}

    
    var body RobotReq
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    if body.SerialNumber == "" {
    http.Error(w, "serial_number is required", http.StatusBadRequest)
    return
}
    newRobot, err := robot.RegisterRobot( hospitalID , body.SerialNumber, body.Model)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newRobot)
}