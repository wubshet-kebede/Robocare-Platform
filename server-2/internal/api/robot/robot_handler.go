package robot

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/robot"
)

type RobotReq struct {
    SerialNumber string `json:"serial_number"`
    Model        string `json:"model"`
}

func CreateRobotHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Get HospitalID (Usually from your Auth Middleware/Context)
    // Assuming you stored it as a string or UUID in context
    hID := r.Context().Value("hospital_id").(uuid.UUID)

    // 2. Decode the JSON body from Nuxt 3
    var body RobotReq
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // 3. Call the Service
    newRobot, err := robot.RegisterRobot(hID, body.SerialNumber, body.Model)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 4. Return the new Robot as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newRobot)
}