package patient

import (
	"fmt"

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
// func GetPatientByRoomID(roomID uuid.UUID) (*model.Patient, error) {
// 	var patient model.Patient
// 	err := db.DB.Where("room_id = ?", roomID).First(&patient).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &patient, nil
// }