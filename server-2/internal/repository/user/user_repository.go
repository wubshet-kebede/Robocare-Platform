package user

import (
	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)
func GetUserByHospitalID(hospitalID uuid.UUID) (*model.User, error) {
	var user model.User
	err := db.DB.Where("hospital_id = ? OR role = ?", hospitalID, "doctor").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}