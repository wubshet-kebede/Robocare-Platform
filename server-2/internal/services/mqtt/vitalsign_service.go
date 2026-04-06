package mqtt

import (
	"fmt"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/admission"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/mqtt"
)

func ProcessRobotVitals(tenantID string, data model.VitalSign) error {
   
    activeAdmission, err := admission.GetActiveAdmissionByID(data.AdmissionID.String())
    if err != nil {
        return fmt.Errorf("admission not found: %v", err)
    }

   
    if activeAdmission.HospitalID.String() != tenantID {
        return fmt.Errorf("security alert: tenant mismatch for admission %s", data.AdmissionID)
    }

   
    data.PatientID = activeAdmission.PatientID
    // data.IsRobotEntry = true 

    
    return mqtt.SaveVitalSign(data)
}