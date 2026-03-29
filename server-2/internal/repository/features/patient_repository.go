package features

import (
	"errors"
	"fmt"

	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
	"gorm.io/gorm"
)

func CreatePatient(patient *model.Patient) (model.Patient, error) {
	
	if err := db.DB.Create(patient).Error; err != nil {
		return model.Patient{}, err
	}
	
	if err := db.DB.Model(&model.Hospital{}).
		Where("id = ?", patient.HospitalID).
		UpdateColumn("total_patients", gorm.Expr("total_patients + ?", 1)).Error; err != nil {
		return model.Patient{}, err
	}
	return *patient, nil
}
func GetAllPatients() ([]model.Patient, error) {
	var patients []model.Patient
	err := db.DB.Find(&patients).Error
	return patients, err
}


func GetPatientByEmailOrPhone(identifier string) (model.Patient, error) {
	var patient model.Patient
	emailRegex := regexp2.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`, 0)
	phoneRegex := regexp2.MustCompile(`^\+?[0-9]{7,15}$`, 0)
	isEmail, _ := emailRegex.MatchString(identifier)
	isPhone, _ := phoneRegex.MatchString(identifier)
	var err error
	switch {
	case isEmail:
		
		err = db.DB.Where("email = ?", identifier).First(&patient).Error
	case isPhone:
		
		err = db.DB.Where("phone = ?", identifier).First(&patient).Error
	default:
		
		return patient, errors.New("invalid identifier format: must be a valid email or phone number")
	}

	return patient, err
}
func CheckForPatients(patientID uuid.UUID) (model.Patient, error) {
	var patient model.Patient
	if err := db.DB.First(&patient, "id = ?", patientID).Error; err != nil {

		return model.Patient{}, err
	}
	return patient, nil
}
func UpdatePatient(patient *model.Patient, hosID uuid.UUID) (model.Patient, error) {
	if hosID == uuid.Nil{
		return model.Patient{}, errors.New("hospital_id is required")
	}
	var existing model.Patient
	if err := db.DB.Where("id = ? AND hospital_id = ?", patient.ID, hosID).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Patient{}, fmt.Errorf("patient not found")
		}
		return model.Patient{}, err
	}
	if err := db.DB.Model(&existing).Updates(patient).Error; err != nil {
		return model.Patient{}, err
	}
	return existing, nil
}
func DeletePatient(id uuid.UUID) (model.Patient, error) {
	var patient model.Patient
	if err := db.DB.First(&patient, "id = ?", id).Error; err != nil {
		return model.Patient{}, err
	}

	if err := db.DB.Delete(&patient).Error; err != nil {
		return model.Patient{}, err
	}
	if err := db.DB.Model(&model.Hospital{}).
		Where("id = ?", patient.HospitalID).
		UpdateColumn("total_patients", gorm.Expr("total_patients - ?", 1)).Error; err != nil {
		return model.Patient{}, err
	}

	return patient, nil
}
func GetPatientByID(id uuid.UUID) (model.Patient, error) {
	var patient model.Patient
	if err := db.DB.First(&patient, "id = ?", id).Error; err != nil {
		return model.Patient{}, err
	}
	return patient, nil
}
func GetPatients(hosID uuid.UUID, page, limit int) ([]model.Patient, int64, error) {
	var patients []model.Patient
	var total int64

	offset := (page - 1) * limit

	if err := db.DB.Model(&model.Patient{}).
		Where("hospital_id = ?", hosID).
		Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}
