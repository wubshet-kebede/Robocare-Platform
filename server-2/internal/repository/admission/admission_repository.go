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

func GetAssignedPatientsRepository(staffID uuid.UUID) ([]model.AssignedPatientResponse, error) {

	var patients []model.AssignedPatientResponse

	err := db.DB.Table("admissions").
	Select(`
		patients.id as patient_id,
		patients.full_name,
		patients.gender,
		patients.date_of_birth,

		admissions.diagnosis,
		admissions.admission_status,
		admissions.urgency,
		admissions.room_id,

		rooms.room_number,

		vs.heart_rate,
		vs.sp_o2,
		vs.temperature,
		vs.measured_at,

		assigned_doctor.full_name as assigned_doctor_name
	`).
	Joins(`
		JOIN patients 
		ON patients.id = admissions.patient_id
	`).
	Joins(`
		LEFT JOIN rooms 
		ON rooms.id = admissions.room_id
	`).
	Joins(`
		LEFT JOIN users as assigned_doctor 
		ON assigned_doctor.id = admissions.assigned_doctor_id
	`).
	Joins(`
		LEFT JOIN LATERAL (
			SELECT *
			FROM vital_signs
			WHERE vital_signs.patient_id = patients.id
			ORDER BY measured_at DESC
			LIMIT 1
		) vs ON true
	`).
	Where(`
		admissions.assigned_doctor_id = ?
		AND admissions.is_active = ?
	`, staffID, true).
	Scan(&patients).Error
	if err != nil {
		return nil, err
	}

	return patients, nil
}
func GetAdmittedPatientVitalsRepository(
	hospitalID uuid.UUID,
	userID uuid.UUID,
) ([]model.AdmittedPatientVitalResponse, error) {

	var vitals []model.AdmittedPatientVitalResponse

	latestVitals := db.DB.
		Table("vital_signs").
		Select("DISTINCT ON (admission_id) *").
		Order("admission_id, measured_at DESC")

	err := db.DB.
		Table("(?) as vs", latestVitals).
		Select(`
			p.id as patient_id,
			a.id as admission_id,
			p.full_name as patient_name,
			a.status,
			vs.heart_rate,
			vs.sp_o2,
			vs.temperature,
			vs.systolic_bp,
			vs.diastolic_bp,
			vs.measured_at
		`).
		Joins("JOIN admissions a ON vs.admission_id = a.id").
		Joins("JOIN patients p ON a.patient_id = p.id").
		Where("a.hospital_id = ?", hospitalID).
		Where("a.assigned_doctor_id = ?", userID).
		Where("a.is_active = ?", true).
		Scan(&vitals).Error

	return vitals, err
}