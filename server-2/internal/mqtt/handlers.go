package mqtt

import (
	"encoding/json"
	"log"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/services/mqtt"
)

func HandleVitalsSubmission(tenantID string, payload []byte) {
    var v model.VitalSign
    if err := json.Unmarshal(payload, &v); err != nil {
        log.Printf("Tenant %s: Error parsing JSON: %v", tenantID, err)
        return
    }
    err := mqtt.ProcessRobotVitals(tenantID, v)
    if err != nil {
        log.Printf("Tenant %s: Failed to save vitals: %v", tenantID, err)
    }
}