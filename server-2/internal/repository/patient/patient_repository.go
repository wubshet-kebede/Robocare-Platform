package patient

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"gorm.io/gorm"
)

func CreatePatient(patient model.Patient) (*model.Patient, error) {
	var existing model.Patient

	
	err := db.DB.Where("phone = ? OR email = ?", patient.Phone, patient.Email).First(&existing).Error
	if err == nil {
		
		return nil, fmt.Errorf("patient with phone '%s' or email '%s' already exists", patient.Phone, patient.Email)
	} else if err != gorm.ErrRecordNotFound {
		
		return nil, err
	}

	
	if err := db.DB.Create(&patient).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}
func GetPatientByID(id uuid.UUID) (*model.Patient, error) {
    var patient model.Patient

    err := db.DB.Where("id = ?", id).First(&patient).Error
    if err != nil {
        return nil, err
    }

    return &patient, nil
}
func GetPatientsWithAdmissions(hospitalID uuid.UUID) ([]model.PatientAdmissionView, error) {

	var result []model.PatientAdmissionView

	err := db.DB.
		Table("patients").
		Select(`
			patients.id as patient_id,
			patients.full_name,
			patients.date_of_birth,
			patients.gender,
			patients.blood_type,
			admissions.status as admission_status,
			admissions.urgency,
			admissions.diagnosis,
			admissions.bed_number,
			patients.Allergies,
			patients.Phone,
			patients.Emergency_contact_name,
			patients.Emergency_contact_phone,
			users.full_name as assigned_doctor_name,
			rooms.room_number
		`).
		Joins(`
			LEFT JOIN admissions
			ON admissions.patient_id = patients.id
			AND admissions.is_active = true
		`).
		Joins(`
			LEFT JOIN users
			ON users.id = admissions.assigned_doctor_id
		`).
		Joins(`
			LEFT JOIN rooms
			ON rooms.id = admissions.room_id
		`).
		Where("patients.hospital_id = ?", hospitalID).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}