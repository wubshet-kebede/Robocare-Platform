package mqttclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/api/ws"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/vitals"
)

// HandleVitalsMessage processes incoming MQTT messages for vitals
func HandleVitalsMessage(wsManager *ws.Manager, msg mqtt.Message) {
    log.Println("Received MQTT message:", msg.Topic())

    hospitalID, robotID, patientID, err := parseVitalsTopic(msg.Topic())
    if err != nil {
        log.Println("Topic parsing error:", err)
        return
    }

    var payload model.VitalSign
    if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
        log.Println("Invalid JSON:", err)
        return
    }

    log.Printf("Vitals → hospital=%s robot=%s patient=%s payload=%s",
        hospitalID, robotID, patientID, string(msg.Payload()))

    
    if payload.PatientID != uuid.Nil && payload.PatientID != patientID {
        log.Println("Patient ID mismatch")
        return
    }
    
    err = vitals.ProcessVitals(hospitalID, robotID, patientID, payload)
    if err != nil {
        log.Println("Failed to process vitals:", err)
    }
   
	wsPayload, _ := json.Marshal(payload)

	wsManager.BroadcastVitals(
		hospitalID.String(),
		wsPayload,
	)
}

func parseVitalsTopic(topic string) (uuid.UUID, uuid.UUID, uuid.UUID, error) {
    parts := strings.Split(topic, "/")
    if len(parts) != 7 || parts[0] != "hospital" || parts[2] != "robot" || parts[4] != "patient" {
    return uuid.Nil, uuid.Nil, uuid.Nil, fmt.Errorf("invalid topic structure")
    }

    hospitalID, err := uuid.Parse(parts[1])
    if err != nil {
        return uuid.Nil, uuid.Nil, uuid.Nil, fmt.Errorf("invalid hospital_id")
    }

    robotID, err := uuid.Parse(parts[3])
    if err != nil {
        return uuid.Nil, uuid.Nil, uuid.Nil, fmt.Errorf("invalid robot_id")
    }

    patientID, err := uuid.Parse(parts[5])
    if err != nil {
        return uuid.Nil, uuid.Nil, uuid.Nil, fmt.Errorf("invalid patient_id")
    }

    return hospitalID, robotID, patientID, nil
}