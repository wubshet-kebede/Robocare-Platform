package admission

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)

func AdmissionRepository(admission model.Admission) (*model.Admission, error) {
    var existingPatient model.Admission
    err := db.DB.Where("patient_id = ? AND is_active = true", admission.PatientID).First(&existingPatient).Error
    if err == nil {
        return nil, fmt.Errorf("patient already has an active admission")
    }
    var existingBed model.Admission
    err = db.DB.Where("room_id = ? AND bed_number = ? AND is_active = true", 
        admission.RoomID, admission.BedNumber).First(&existingBed).Error
    if err == nil {
        return nil, fmt.Errorf("bed %s in this room is already occupied", admission.BedNumber)
    }
    if err := db.DB.Create(&admission).Error; err != nil {
        return nil, err
    }
    return &admission, nil
}
func GetActiveAdmissionByID(admissionID string) (*model.Admission, error) {
    var admission model.Admission
    err := db.DB.Where("id = ? AND is_active = true", admissionID).First(&admission).Error
    if err != nil {
        return nil, fmt.Errorf("active admission not found: %v", err)
    }
    return &admission, nil
}
func GetPatientByRoomID(roomID uuid.UUID) (*model.Admission, error) {
	var admission model.Admission
	err := db.DB.Where("room_id = ? OR is_active = true", roomID).First(&admission).Error
	if err != nil {
		return nil, err
	}
	return &admission, nil
}
func GetPatientAdmissionID(patientID uuid.UUID) (*model.Admission, error) {
    var admission model.Admission
    err := db.DB.Where("patient_id = ? AND is_active = true", patientID).First(&admission).Error
    if err != nil {
        return nil, fmt.Errorf("active admission not found for patient: %v", err)
    }
    return &admission, nil
}