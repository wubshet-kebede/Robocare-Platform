package admission

import (
	"fmt"

	"errors"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"gorm.io/gorm"
)

func AdmissionRepository(admission model.Admission) (*model.Admission, error) {

	var patient model.Patient

	err := db.DB.
		Where("id = ? AND hospital_id = ?",
			admission.PatientID,
			admission.HospitalID,
		).
		First(&patient).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("patient not found")
		}

		return nil, err
	}
	var existingAdmission model.Admission

	err = db.DB.
		Where("patient_id = ? AND is_active = ?",
			admission.PatientID,
			true,
		).
		First(&existingAdmission).Error

	if err == nil {
		return nil, fmt.Errorf("patient already has an active admission")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var occupiedBed model.Admission

	err = db.DB.
		Where(
			"room_id = ? AND bed_number = ? AND is_active = ?",
			admission.RoomID,
			admission.BedNumber,
			true,
		).
		First(&occupiedBed).Error

	if err == nil {
		return nil, fmt.Errorf(
			"bed %s is already occupied",
			admission.BedNumber,
		)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
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
// repository/admission.go


