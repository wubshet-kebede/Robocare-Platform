package organization

import (
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)


func GetHospitalByHospitalID(hospitalID uuid.UUID) (*model.Hospital, error) {
	var hospital model.Hospital
	err := db.DB.Where("id = ? ", hospitalID).First(&hospital).Error
	if err != nil {
		return nil, err
	}
	return &hospital, nil
}