package patient

import (
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/repository/patient"
)
func CreatePatient(newPatient model.Patient) (*model.Patient, error){
	patient, err := patient.CreatePatient(newPatient)
	return patient, err
}
func GetPatients(hospitalID uuid.UUID) ([]model.Patient, error) {
	patients, err := patient.GetPatients(hospitalID)
	return patients, err
}