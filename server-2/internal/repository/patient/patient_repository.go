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
func GetPatientsWithAdmissions(hospitalID uuid.UUID) ([]model.PatientWithAdmission, error) {
	var patients []model.Patient

	err := db.DB.
		Where("hospital_id = ?", hospitalID).
		Find(&patients).Error

	if err != nil {
		return nil, err
	}

	var response []model.PatientWithAdmission

	for _, patient := range patients {

		var admission model.Admission

		err := db.DB.
			Where("patient_id = ? AND is_active = ?", patient.ID, true).
			First(&admission).Error

		if err != nil {

			response = append(response, model.PatientWithAdmission{
				Patient:         patient,
				ActiveAdmission: nil,
			})

			continue
		}

		response = append(response, model.PatientWithAdmission{
			Patient:         patient,
			ActiveAdmission: &admission,
		})
	}

	return response, nil
}