package vitals

import (
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)
func SaveVitals(hospitalID,admissionID, robotID, patientID uuid.UUID, vital model.VitalSign) error {

    vital.HospitalID = hospitalID
    vital.AdmissionID = admissionID
    vital.RobotID = robotID
    vital.PatientID = patientID
    vital.IsRobotEntry = true

    return db.DB.Create(&vital).Error
}
