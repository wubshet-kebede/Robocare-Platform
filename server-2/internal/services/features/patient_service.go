package features

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/features"
)

func CreatePatient(patient model.Patient) (model.Patient, error) {
	// hashed, err := utils.HashPassword(password)
	// if err != nil {
	// 	return model.Patient{}, err
	// }
	// patient.PasswordHash = hashed
	createdPatient, err := features.CreatePatient(&patient)

	if err != nil {
		return model.Patient{}, err
	}
	return createdPatient, nil
}
func ListPatients() ([]model.Patient, error) {
	return features.GetAllPatients()
}
func CheckForPatient(patientID uuid.UUID) (model.Patient, error) {
	return features.CheckForPatients(patientID)
}
func UpdatePatient(patient model.Patient, hosID uuid.UUID) (model.Patient, error) {
	updatedPatient, err := features.UpdatePatient(&patient, hosID)

	if err != nil {
		return model.Patient{}, err
	}
	return updatedPatient, nil
}
func DeletePatient(id uuid.UUID, patientHosID uuid.UUID) (model.Patient, error) {
	patient, err := features.GetPatientByID(id)
	if err != nil {
		return model.Patient{}, err
	}
	if patient.HospitalID != patientHosID {
		return model.Patient{}, fmt.Errorf("forbidden: cannot delete patient from another hospital")
	}
	deletedPatient, err := features.DeletePatient(id)

	if err != nil {
		return model.Patient{}, err
	}
	return deletedPatient, nil
}

func GetPatient(patientID uuid.UUID, patientHosID uuid.UUID) (model.Patient, error) {

	patient, err := features.GetPatientByID(patientID)
	if err != nil {
		return model.Patient{}, err
	}

	if patient.HospitalID != patientHosID {
		return model.Patient{}, fmt.Errorf("forbidden: cannot access patient from another hospital")
	}

	return patient, nil
}

func GetPatients(hosID uuid.UUID, page, limit int) ([]model.Patient, int64, error) {
	patients, total, err := features.GetPatients(hosID, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return patients, total, nil
}
