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
func  GetUserWithHospital(userID uuid.UUID) (*model.User, *model.Hospital, error) {
	var user model.User
	var hospital model.Hospital
	if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, nil, err
	}
	if err := db.DB.First(&hospital, "id = ?", user.HospitalID).Error; err != nil {
		return nil, nil, err
	}

	return &user, &hospital, nil
}